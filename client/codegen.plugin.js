export function plugin(schema, documents, config, info) {
  const typesMap = schema.getTypeMap();

  return Object.keys(typesMap).join('\n');
}
