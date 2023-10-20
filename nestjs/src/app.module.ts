import * as dotenv from 'dotenv';
dotenv.config();
import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { BankAccountsModule } from './bank-accounts/bank-accounts.module';
import { TypeOrmModule } from '@nestjs/typeorm';
import { BankAccount } from './bank-accounts/entities/bank-account.entity';
import { PixKeysModule } from './pix-keys/pix-keys.module';
import { PixKey } from './pix-keys/entities/pix-key.entity';

@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: 'postgres',
      host: process.env.POSTGRES_HOST,
      database: 'nest',
      username: 'postgres',
      password: process.env.POSTGRES_PASSWORD,
      entities: [BankAccount, PixKey],
      synchronize: true,
    }),
    BankAccountsModule,
    PixKeysModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule { }
