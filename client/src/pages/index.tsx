import Head from 'next/head';

import { Header } from 'src/components/Header';
import { PostList } from 'src/components/PostList';

const Home = () => {
  return (
    <div className="relative h-screen bg-gradient-to-b from-gray-900/10 to-[#010511] lg:h-[140vh]">
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
