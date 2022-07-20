import Image from 'next/image';
import { AiOutlineHeart } from 'react-icons/ai';
import { FaRegComment } from 'react-icons/fa';

import { GetAllPostsQuery } from '../generated/graphql';

export const Post = ({ post }: { post: GetAllPostsQuery['getAllPosts'] }) => {
  return (
    <button>
      <div className="card w-96 bg-base-100 shadow-xl">
        <Image
          className="bg-fixed"
          src="https://api.lorem.space/image/shoes?w=400&h=225"
          width={400}
          height={225}
          layout="responsive"
          alt="Shoes"
        />
        <div className="card-body">
          <h2 className="card-title">
            {post[0]?.title}
            <div className="badge badge-secondary">NEW</div>
          </h2>
          <div className="card-actions justify-end">
            <div className="badge badge-outline">
              <AiOutlineHeart />
              {post[0]?.comment.length}
            </div>
            <div className="badge badge-outline">
              <FaRegComment />
              {post[0]?.like.length}
            </div>
          </div>
        </div>
      </div>
    </button>
  );
};
