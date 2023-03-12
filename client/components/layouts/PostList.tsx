import { useQuery } from 'urql';

import { PostsDocument } from 'queries/query.generated';

export const PostList = (props) => {
  const [res] = useQuery({ query: PostsDocument });
  return (
    <div>
      {res.data?.posts.map((post) => (
        <div key={post.id}>
          {post.id} - {post.body}
        </div>
      ))}
    </div>
  );
};
