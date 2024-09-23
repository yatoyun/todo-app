import {
  generateSchemaTypes,
  generateReactQueryComponents,
} from "@openapi-codegen/typescript";
import { defineConfig } from "@openapi-codegen/cli";
export default defineConfig({
  // ここのprojectNameを自動生成時にkeyとして用いる
  todoApp: {
    from: {
      // openAPIのパス
      relativePath: "../openapi/swagger.yaml",
      source: "file",
    },
    // 出力先のパス
    outputDir: "./src/openapi",
    to: async (context) => {
      // ESLintルール違反コードを大量に生成するため、eslint-disableを自動で付加
      // 参照: https://github.com/fabien0102/openapi-codegen/issues/96#issuecomment-1243867098
      const writeFile = context.writeFile;
      context.writeFile = (file, data) =>
        writeFile(file, `/* eslint-disable */\n\n${data}`);
      const filenamePrefix = "todoApp";
      const { schemasFiles } = await generateSchemaTypes(context, {
        filenamePrefix,
      });
      await generateReactQueryComponents(context, {
        filenamePrefix,
        schemasFiles,
      });
    },
  },
});
