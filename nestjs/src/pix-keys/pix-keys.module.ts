import { Module } from '@nestjs/common';
import { PixKeysService } from './pix-keys.service';
import { PixKeysController } from './pix-keys.controller';

@Module({
  controllers: [PixKeysController],
  providers: [PixKeysService],
})
export class PixKeysModule {}
