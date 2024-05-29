import { Injectable } from '@nestjs/common';
import { Request, Config } from './api';

@Injectable()
export class AppService {
  getHello(): string {
    return 'Hello World!';
  }
  convert(req: Request): Config {
    return req.config;
  }
}
