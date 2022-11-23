import { usePostsQuery } from 'app/generated/graphql';

import { PostCard } from './components/PostCard';

export const PostPage = () => {
  const [result] = usePostsQuery();
  const { data, fetching, error } = result;

  if (fetching) {
    return <div>fetching...</div>;
  }
  if (error) {
    return <div>{error.message}</div>;
  }
  if (!data) {
    return <div>No Data</div>;
  }

  return (
    <div className="flex justify-center p-10">
      <div className="grid grid-cols-3 gap-10 ">
        {data?.posts.map((post) => (
          <PostCard key={post.id}>{post}</PostCard>
        ))}
      </div>
    </div>
  );
};
