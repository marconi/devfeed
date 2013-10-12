({
  baseUrl: ".",
  name: "vendors/almond",
  include: "main",
  mainConfigFile: "main.js",
  insertRequire: ['main'],
  out: "main.built.js",
  findNestedDependencies: true
})