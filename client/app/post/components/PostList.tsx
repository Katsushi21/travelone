import { PostsDocument } from '../../generated/graphql';
import { graphQLClient } from '../../GraphQLRequestClient';
import { PostCard } from './PostCard';

export const PostList = async () => {
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
