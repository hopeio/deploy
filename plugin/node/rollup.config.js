import resolve from '@rollup/plugin-node-resolve';
import typescript from '@rollup/plugin-typescript';
import commonjs from '@rollup/plugin-commonjs';
import json from '@rollup/plugin-json';
import glob from 'fast-glob';
import terser from '@rollup/plugin-terser';
import alias from '@rollup/plugin-alias';
import { fileURLToPath } from 'node:url';
import { readFileSync } from 'fs';
import path from 'node:path';

const __filename = fileURLToPath(import.meta.url);
const packageJson = JSON.parse(readFileSync('./package.json', 'utf8')); // 读取UMD全局模块名，在package中定义了
const pkgName = packageJson.name;
const __dirname = path.dirname(__filename);

const pathResolve = (p) => path.resolve(__dirname, p);
const input = Object.fromEntries(
  glob
    .sync('src/**/*.(ts|js)', {
      ignore: [`src/**/*.(test|spec).(ts|js)`],
    })
    .map((file) => [
      // 这里将删除 `src/` 以及每个文件的扩展名。
      // 因此，例如 src/nested/foo.js 会变成 nested/foo
      path.relative(
        'src',
        file.slice(0, file.length - path.extname(file).length),
      ),
      // 这里可以将相对路径扩展为绝对路径，例如
      // src/nested/foo 会变成 /project/src/nested/foo.js
      fileURLToPath(new URL(file, import.meta.url)),
    ]),
);

export default {
  input: input,
  output: {
    dir: 'dist',
    format: 'esm',
  },
  plugins: [
    typescript({
      tsconfig: './tsconfig.json',
    }),
    resolve(),
    alias({
      entries: {
        '@': pathResolve('src'),
        _: __dirname,
      },
    }),
    commonjs({
      include: /node_modules/,
    }),
    json(),
    terser(),
  ],
};
