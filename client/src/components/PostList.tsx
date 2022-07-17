import { GetServerSideProps } from 'next';

import { useGetAllPostsQuery } from 'src/generated/graphql';
import { Post } from './Post';

import { Post } from './Post';

export const PostList = ({ resultPosts }) => {
  console.log(resultPosts);
  return (
    <div>
      <Post />
    </div>
  );
};

export const getServerSideProps: GetServerSideProps = async () => {
  const [resultPosts] = useGetAllPostsQuery();
  return {
    props: resultPosts,
  };
};
