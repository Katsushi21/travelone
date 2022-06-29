import React from 'react';

import { useGetAllPostsQuery } from 'src/generated/graphql';
import { Post } from './Post';

export const PostList: React.FC = () => {
  const [result] = useGetAllPostsQuery();
  return (
    <div>
      <Post />
    </div>
  );
};
