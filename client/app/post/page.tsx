import { sdk } from '../GraphQLRequestClient';
import { PostCard } from './components/PostCard';

const PostPage = async () => {
  const data = await sdk.Posts();

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
