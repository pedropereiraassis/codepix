import { Injectable } from '@nestjs/common';
import { CreatePixKeyDto } from './dto/create-pix-key.dto';
import { UpdatePixKeyDto } from './dto/update-pix-key.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { PixKey } from './entities/pix-key.entity';
import { Repository } from 'typeorm';
import { BankAccount } from 'src/bank-accounts/entities/bank-account.entity';

@Injectable()
export class PixKeysService {

  constructor(
    @InjectRepository(PixKey) private pixKeyRepo: Repository<PixKey>,
    @InjectRepository(BankAccount) private bankAccountRepo: Repository<BankAccount>,
  ) { }

  async create(bankAccountId: string, createPixKeyDto: CreatePixKeyDto) {
    await this.bankAccountRepo.findOneOrFail({
      where: { id: bankAccountId },
    });

    return await this.pixKeyRepo.save({
      bank_account_id: bankAccountId,
      ...createPixKeyDto,
    })
  }

  async findAll(bankAccountId: string) {
    return await this.pixKeyRepo.find({
      where: { bank_account_id: bankAccountId },
      order: { created_at: 'DESC' },
    });
  }
}
