import { AiOutlineHeart } from 'react-icons/ai';
import { FaRegComment } from 'react-icons/fa';

import { PostsQuery } from '../../generated/graphql';

type Props = {
  key: string;
  post: PostsQuery['posts'][number];
};

export const PostCard = (props: Props) => {
  const post = props.post;

  const likeNum = post.likes ? post.likes.length : 0;
  const commentNum = post.comments ? post.comments.length : 0;

  return (
    <div className="card w-96 bg-base-100 shadow-xl">
      {/* <Image
          className="bg-fixed"
          src={post.img}
          width={400}
          height={225}
          alt="*"
        /> */}
      <div className="card-body">
        <h2 className="card-title">{post.title}</h2>
        <div className="flex items-center justify-between">
          <div className="flex items-center font-medium">
            <label
              tabIndex={0}
              className="btn-ghost btn-circle avatar btn mx-2"
            >
              <div className="w-10 rounded-full">
                {/* <Image
                    src={post.account.avatar}
                    width={10}
                    height={10}
                    alt="*"
                  /> */}
              </div>
            </label>
            {post.account.name}
          </div>
          <div>
            <div className="badge badge-outline mx-2 h-6 w-14 text-xl">
              <AiOutlineHeart />
              {likeNum}
            </div>
            <div className="badge badge-outline mx-2 h-6 w-14 text-xl">
              <FaRegComment />
              {commentNum}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
