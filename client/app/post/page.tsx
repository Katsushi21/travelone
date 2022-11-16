'use client';

import { usePostsQuery } from 'app/generated/graphql';

import { PostCard } from './components/PostCard';

const PostPage = () => {
  const [resultPosts] = usePostsQuery();
  const { data, fetching, error } = resultPosts;
  const postToRender = data?.posts;

  return (
    <div className="flex justify-center p-10">
      <div className="grid grid-cols-3 gap-10 ">
        {postToRender?.map((post) => (
          <PostCard key={post?.id} post={post} />
        ))}
      </div>
    </div>
  );
};

export default PostPage;
