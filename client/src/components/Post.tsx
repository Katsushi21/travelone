import Image from 'next/image';
import { AiOutlineHeart } from 'react-icons/ai';
import { FaRegComment } from 'react-icons/fa';

import { AllPostsQuery } from '../generated/graphql';

export const Post = ({ post }: { post: AllPostsQuery['allPosts'][number] }) => {
  return (
    <button>
      <div className="card w-96 bg-base-100 shadow-xl">
        <Image
          className="bg-fixed"
          src={post.img}
          width={400}
          height={225}
          layout="responsive"
          alt="*"
        />
        <div className="card-body">
          <h2 className="card-title">{post.title}</h2>
          <div className="flex items-center justify-between">
            <div className="flex items-center font-medium">
              <label
                tabIndex={0}
                className="btn-ghost btn-circle avatar btn mx-2"
              >
                <div className="w-10 rounded-full">
                  <Image
                    src={post.account.avatar}
                    width={10}
                    height={10}
                    layout="responsive"
                    alt="*"
                  />
                </div>
              </label>
              {post.account.name}
            </div>
            <div>
              <div className="badge badge-outline mx-2 h-6 w-14 text-xl">
                <AiOutlineHeart />
                {post.like.length}
              </div>
              <div className="badge badge-outline mx-2 h-6 w-14 text-xl">
                <FaRegComment />
                {post.comment.length}
              </div>
            </div>
          </div>
        </div>
      </div>
    </button>
  );
};
