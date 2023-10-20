import { Injectable } from '@nestjs/common';
import { CreateBankAccountDto } from './dto/create-bank-account.dto';
import { Repository } from 'typeorm';
import { BankAccount } from './entities/bank-account.entity';
import { InjectRepository } from '@nestjs/typeorm';

@Injectable()
export class BankAccountsService {
  constructor(
    @InjectRepository(BankAccount)
    private bankAccountRepo: Repository<BankAccount>,
  ) {}

  // DTO - Data Transfer Object
  async create(createBankAccountDto: CreateBankAccountDto) {
    return await this.bankAccountRepo.save(createBankAccountDto);
  }

  async findAll() {
    return await this.bankAccountRepo.find();
  }

  async findOne(id: string) {
    return await this.bankAccountRepo.findOneOrFail({
      where: { id },
    });
  }
}
