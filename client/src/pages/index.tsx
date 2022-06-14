import Head from 'next/head';
import Image from 'next/image';

import { Header } from 'src/components/Header';
import { PostList } from 'src/components/PostList';

const Home = () => {
  return (
    <div className="relative h-screen bg-white from-gray-900/10 to-[#010511] lg:h-[140vh]">
      <Image
        src="/background.jpg"
        width={4460}
        height={3345}
        layout="responsive"
        alt="background image"
      />
      <Head>
        <title>Traveling</title>
      </Head>
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
      {/* Footer */}
    </div>
  );
};

export default Home;
