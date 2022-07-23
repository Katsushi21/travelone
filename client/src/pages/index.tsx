import Image from 'next/image';
import Script from 'next/script';

import { Footer } from '@/components/Footer';
import { Header } from 'src/components/Header';
import { PostList } from 'src/components/PostList';

const Home = () => {
  return (
    <>
      <Script src="../../.node_modules/tw-elements-ELEMENTS-PATH/dist/js/index.min.js" />
      <div data-theme="light">
        <Header />
        <Image
          className="bg-fixed"
          src="/background.jpg"
          width={4460}
          height={3345}
          layout="responsive"
          alt="background image"
        />
        <main>
          {/* Banner */}
          <section>
            <PostList />
            {/* Row */}
            {/* Row */}
            {/* Row */}
            {/* Row */}
            {/* Row */}
            {/* Row */}
            {/* Row */}
          </section>
        </main>
        {/* Modal */}
        <Footer />
      </div>
    </>
  );
};

export default Home;
