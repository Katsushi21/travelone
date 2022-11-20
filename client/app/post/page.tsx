'use client';

import { useState } from 'react';

import { usePostsQuery } from 'app/generated/graphql';

import { PostCard } from './components/PostCard';

const PostPage = () => {
  const [posts, setPosts] = useState({});
  const [page, setPage] = useState(1);

  // useEffect(() => {
  //   first;

  //   return () => {
  //     second;
  //   };
  // }, [page]);

  const [resultPosts] = usePostsQuery();
  const { data, fetching, error } = resultPosts;

  if (fetching) return <p>Loading...</p>;
  if (error) return <p>Oh no... {error.message}</p>;

  const postsToRender = data?.posts;

  return (
    <div className="flex justify-center p-10">
      <div className="grid grid-cols-3 gap-10 ">
        {postsToRender?.map((post) => (
          <PostCard key={post?.id} post={post} />
        ))}
      </div>
    </div>
  );
};

export default PostPage;
