import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { CreatePixKeyDto } from './dto/create-pix-key.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { PixKey, PixKeyKind } from './entities/pix-key.entity';
import { Repository } from 'typeorm';
import { BankAccount } from 'src/bank-accounts/entities/bank-account.entity';
import { ClientGrpc } from '@nestjs/microservices';
import { PixKeyClientGrpc, RegisterPixKeyRpcResponse } from './pix-keys.grpc';
import { lastValueFrom } from 'rxjs';

@Injectable()
export class PixKeysService implements OnModuleInit {

  private pixGrpcService: PixKeyClientGrpc;

  constructor(
    @InjectRepository(PixKey) private pixKeyRepo: Repository<PixKey>,
    @InjectRepository(BankAccount) private bankAccountRepo: Repository<BankAccount>,
    @Inject('PIX_PACKAGE') private pixGrpcPackage: ClientGrpc,
  ) { }

  onModuleInit() {
    this.pixGrpcService = this.pixGrpcPackage.getService('PixService');
  }

  async create(bankAccountId: string, createPixKeyDto: CreatePixKeyDto) {
    await this.bankAccountRepo.findOneOrFail({
      where: { id: bankAccountId },
    });

    const remotePixKey = await this.findRemotePixKey(createPixKeyDto);

    return remotePixKey
      ? await this.createIfNotExists(bankAccountId, remotePixKey)
      : await this.pixKeyRepo.save({ bank_account_id: bankAccountId, ...createPixKeyDto })
  }

  private async findRemotePixKey(data: { key: string; kind: string }): Promise<RegisterPixKeyRpcResponse | null> {
    try {
      return await lastValueFrom(this.pixGrpcService.find(data));
    } catch (err) {
      console.error(err);
      if (err.details == 'no key was found') {
        return null;
      }

      throw new PixKeyGrpcUnknownError('Grpc Internal Eerror');
    }
  }

  private async createIfNotExists(bankAccountId: string, remotePixKey: RegisterPixKeyRpcResponse) {
    const hasLocalPixKey = await this.pixKeyRepo.exist({
      where: {
        key: remotePixKey.key,
      },
    });

    if (hasLocalPixKey) {
      throw new PixKeyAlreadyExistsError();
    } else {
      return this.pixKeyRepo.save({
        id: remotePixKey.id,
        bank_account_id: bankAccountId,
        key: remotePixKey.key,
        kind: remotePixKey.kind as PixKeyKind,
      });
    }
  }

  async findAll(bankAccountId: string) {
    return await this.pixKeyRepo.find({
      where: { bank_account_id: bankAccountId },
      order: { created_at: 'DESC' },
    });
  }
}

export class PixKeyGrpcUnknownError extends Error { }
export class PixKeyAlreadyExistsError extends Error { }