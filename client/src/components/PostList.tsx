import React from 'react';

import { useGetAllPostsQuery } from 'src/generated/graphql';

export const PostList: React.FC = () => {
  const [result] = useGetAllPostsQuery();
  console.log(result);
  return <div>PostList</div>;
};
