import { useAllPostsQuery } from 'src/generated/graphql';

import { Post } from './Post';

export const PostList = () => {
  const [resultPosts] = useAllPostsQuery();

  const { data, fetching, error } = resultPosts;

  if (fetching) return <div>Fetching</div>;
  if (error) return <div>Error!</div>;

  const postToRender = data?.allPosts;

  return (
    <div className="flex justify-center p-10">
      <div className="grid grid-cols-3 gap-10 ">
        {postToRender?.map((post) => (
          <Post key={post?.id} post={post} />
        ))}
      </div>
    </div>
  );
};
