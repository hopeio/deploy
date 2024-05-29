import { Controller, Get, Post, Body } from '@nestjs/common';
import { AppService } from './app.service';
import { Request, Config } from './api';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }

  @Post()
  convert(@Body() config: Request): Config {
    return this.appService.convert(config);
  }
}
