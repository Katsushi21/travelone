'use client';
import { usePostsQuery } from 'app/generated/graphql';

import { useGraphQLClient } from '../GraphQLClientProvider';
import { PostCard } from './components/PostCard';

const PostPage = () => {
  const client = useGraphQLClient();
  const { status, data, error, isFetching } = usePostsQuery(client, {});

  if (isFetching) {
    return <div>fetching...</div>;
  }
  if (!data) {
    return <div>No Data</div>;
  }

  return (
    <div className="flex justify-center p-10">
      <div className="grid grid-cols-3 gap-10 ">
        {data.posts.map((post) => (
          <PostCard key={post.id} post={post} />
        ))}
      </div>
    </div>
  );
};

export default PostPage;
