import Image from 'next/image';

import { Footer } from '../components/Footer';
import { Header } from '../components/Header';

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
      <body>
        <div data-theme="light">
          <header>
            <Header />
          </header>
          <Image
            className="bg-fixed"
            src="/background.jpg"
            width={4460}
            height={3345}
            layout="responsive"
            alt="background image"
          />
          <main>{children}</main>
        </div>
        <Footer />
      </body>
    </html>
  );
}
