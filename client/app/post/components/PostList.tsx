import { use } from 'react';

import { PostsDocument, PostsQuery } from 'app/generated/graphql';

import { graphQLClient } from '../../GraphQLRequestClient';
import { PostCard } from './PostCard';

const getPosts = async (): Promise<PostsQuery> => {
  return await graphQLClient.request(PostsDocument, {});
};

export const PostList = () => {
  const data = use(getPosts());
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
