export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ja">
      <head>
        <title>travelone</title>
      </head>
      <body>{children}</body>
    </html>
  );
}
