import Image from 'next/image';

import { Footer } from '@/components/Footer';
import { Header } from 'src/components/Header';
import { PostList } from 'src/components/PostList';

const Home = () => {
  return (
    <div data-theme="light">
      <Image
        className="bg-fixed"
        src="/background.jpg"
        width={4460}
        height={2375}
        layout="responsive"
        alt="background image"
      />
      <Header />
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
  );
};

export default Home;
