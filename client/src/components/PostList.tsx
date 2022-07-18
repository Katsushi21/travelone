import { GetServerSideProps, NextPage } from 'next';

import { useGetAllPostsQuery } from 'src/generated/graphql';

import { Post } from './Post';

type Props = {
  aaa: string;
  bbb: string;
};

export const PostList: NextPage<Props> = () => {
  const [resultPosts] = useGetAllPostsQuery();
  const { data, fetching, error } = resultPosts;

  if (fetching) return <div>Fetching</div>;
  if (error) return <div>Error!</div>;

  const postToRender = data?.getAllPosts;

  return (
    <div>
      <div className="grid grid-cols-4 gap-20">
        {postToRender?.map((post) => (
          <Post key={post?.id} post={[post]} />
        ))}
      </div>
    </div>
  );
};

// export const getServerSideProps: GetServerSideProps = async () => {
//   const [resultPosts] = useGetAllPostsQuery();
//   return {
//     props: resultPosts,
//   };
// };
