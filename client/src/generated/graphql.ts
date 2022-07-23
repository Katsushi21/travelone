import gql from 'graphql-tag';
import * as Urql from 'urql';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
  Upload: any;
};

export type Account = {
  __typename?: 'Account';
  age: Scalars['Int'];
  avatar: Scalars['String'];
  comment: Array<Maybe<Comment>>;
  createdAt: Scalars['Time'];
  email: Scalars['String'];
  friend: Array<Maybe<Friend>>;
  gender: AccountGender;
  id: Scalars['ID'];
  introduction: Scalars['String'];
  like: Array<Maybe<Like>>;
  mute: Array<Maybe<Mute>>;
  name: Scalars['String'];
  password: Scalars['String'];
  post: Array<Maybe<Post>>;
  type: AccountType;
  updatedAt: Scalars['Time'];
};

export enum AccountGender {
  Female = 'female',
  Male = 'male',
  None = 'none',
}

export type AccountInput = {
  age: Scalars['Int'];
  avatar: Scalars['String'];
  email: Scalars['String'];
  gender: AccountGender;
  introduction: Scalars['String'];
  name: Scalars['String'];
  password: Scalars['String'];
  type: AccountType;
};

export enum AccountType {
  Active = 'active',
  Admin = 'admin',
  Inactive = 'inactive',
}

export type Comment = {
  __typename?: 'Comment';
  account: Account;
  accountId: Scalars['ID'];
  body: Scalars['String'];
  createdAt: Scalars['Time'];
  id: Scalars['ID'];
  post: Post;
  postId: Scalars['ID'];
  updatedAt: Scalars['Time'];
};

export type CommentInput = {
  accountId: Scalars['ID'];
  body: Scalars['String'];
  postId: Scalars['ID'];
};

export type File = {
  __typename?: 'File';
  path: Scalars['String'];
};

export type Friend = {
  __typename?: 'Friend';
  accountId: Scalars['ID'];
  friend: Account;
  friendId: Scalars['ID'];
};

export type FriendInput = {
  accountId: Scalars['ID'];
  friendId: Scalars['ID'];
};

export type Like = {
  __typename?: 'Like';
  account: Account;
  accountId: Scalars['ID'];
  createdAt: Scalars['Time'];
  post: Post;
  postId: Scalars['ID'];
};

export type LikeInput = {
  accountId: Scalars['ID'];
  postId: Scalars['ID'];
};

export type LoginInput = {
  email: Scalars['String'];
  password: Scalars['String'];
};

export type Marker = {
  __typename?: 'Marker';
  createdAt: Scalars['Time'];
  id: Scalars['ID'];
  lat: Scalars['String'];
  lng: Scalars['String'];
  post: Post;
  postId: Scalars['ID'];
  title: Scalars['String'];
  updatedAt: Scalars['Time'];
};

export type MarkerInput = {
  lat: Scalars['String'];
  lng: Scalars['String'];
  postId?: InputMaybe<Scalars['ID']>;
  title: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createAccount: Account;
  createComment: Comment;
  createFriend: Friend;
  createLike: Like;
  createMarker: Marker;
  createMute: Mute;
  createPost: Post;
  createRequest: Request;
  createSession: Session;
  deleteAccount: Account;
  deleteComment: Comment;
  deleteFriend: Friend;
  deleteLike: Like;
  deleteMarker: Marker;
  deleteMute: Mute;
  deletePost: Post;
  deleteRequest: Request;
  deleteSession: Session;
  loginQuery: Account;
  updateAccount: Account;
  updateMarker: Marker;
  updatePost: Post;
  updateRequest: Request;
  updateSession: Session;
  uploadFile: File;
};

export type MutationCreateAccountArgs = {
  input: AccountInput;
};

export type MutationCreateCommentArgs = {
  input: CommentInput;
};

export type MutationCreateFriendArgs = {
  input: FriendInput;
};

export type MutationCreateLikeArgs = {
  input: LikeInput;
};

export type MutationCreateMarkerArgs = {
  input: MarkerInput;
};

export type MutationCreateMuteArgs = {
  input: MuteInput;
};

export type MutationCreatePostArgs = {
  input: PostInput;
};

export type MutationCreateRequestArgs = {
  input: RequestInput;
};

export type MutationCreateSessionArgs = {
  input: SessionInput;
};

export type MutationDeleteAccountArgs = {
  id: Scalars['ID'];
};

export type MutationDeleteCommentArgs = {
  id: Scalars['ID'];
};

export type MutationDeleteFriendArgs = {
  input: FriendInput;
};

export type MutationDeleteLikeArgs = {
  input: LikeInput;
};

export type MutationDeleteMarkerArgs = {
  id: Scalars['ID'];
};

export type MutationDeleteMuteArgs = {
  input?: InputMaybe<MuteInput>;
};

export type MutationDeletePostArgs = {
  id: Scalars['ID'];
};

export type MutationDeleteRequestArgs = {
  input: RequestInput;
};

export type MutationDeleteSessionArgs = {
  input: SessionInput;
};

export type MutationLoginQueryArgs = {
  input: LoginInput;
};

export type MutationUpdateAccountArgs = {
  id: Scalars['ID'];
  input: AccountInput;
};

export type MutationUpdateMarkerArgs = {
  id: Scalars['ID'];
  input: MarkerInput;
};

export type MutationUpdatePostArgs = {
  id: Scalars['ID'];
  input: PostInput;
};

export type MutationUpdateRequestArgs = {
  input: RequestInput;
};

export type MutationUpdateSessionArgs = {
  input: SessionInput;
};

export type MutationUploadFileArgs = {
  input: UploadFile;
};

export type Mute = {
  __typename?: 'Mute';
  accountId: Scalars['ID'];
  mute: Account;
  muteId: Scalars['ID'];
};

export type MuteInput = {
  accountId: Scalars['ID'];
  muteId: Scalars['ID'];
};

export type Post = {
  __typename?: 'Post';
  account: Account;
  accountId: Scalars['ID'];
  body: Scalars['String'];
  comment: Array<Maybe<Comment>>;
  createdAt: Scalars['Time'];
  id: Scalars['ID'];
  img: Scalars['String'];
  like: Array<Maybe<Like>>;
  marker?: Maybe<Marker>;
  title: Scalars['String'];
  updatedAt: Scalars['Time'];
};

export type PostInput = {
  accountId: Scalars['ID'];
  body: Scalars['String'];
  img: Scalars['String'];
  title: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  accountPageInfo: Account;
  allMarkers: Array<Marker>;
  allPosts: Array<Post>;
  likesByPost: Array<Like>;
  myPageInfo: Account;
  requestsByAccountId: Array<Request>;
  requestsByTargetId: Array<Request>;
  sessionByAccountId: Session;
};

export type QueryAccountPageInfoArgs = {
  id: Scalars['ID'];
};

export type QueryLikesByPostArgs = {
  postId: Scalars['ID'];
};

export type QueryMyPageInfoArgs = {
  id: Scalars['ID'];
};

export type QueryRequestsByAccountIdArgs = {
  accountId: Scalars['ID'];
};

export type QueryRequestsByTargetIdArgs = {
  targetAccountId: Scalars['ID'];
};

export type QuerySessionByAccountIdArgs = {
  accountId: Scalars['ID'];
};

export type Request = {
  __typename?: 'Request';
  account: Account;
  accountId: Scalars['ID'];
  status: RequestStatus;
  targetAccount: Account;
  targetAccountId: Scalars['ID'];
};

export type RequestInput = {
  accountId: Scalars['ID'];
  status?: InputMaybe<RequestStatus>;
  targetAccountId: Scalars['ID'];
};

export enum RequestStatus {
  Accept = 'accept',
  BreakAccept = 'breakAccept',
  BreakDeny = 'breakDeny',
  BreakInProcess = 'breakInProcess',
  Cancel = 'cancel',
  Deny = 'deny',
  InProcess = 'inProcess',
}

export type Session = {
  __typename?: 'Session';
  accountId: Scalars['ID'];
  session: Scalars['String'];
  updatedAt: Scalars['Time'];
};

export type SessionInput = {
  accountId: Scalars['ID'];
};

export type UploadFile = {
  content: Scalars['Upload'];
};

export type CreateAccountMutationVariables = Exact<{
  email: Scalars['String'];
  password: Scalars['String'];
  type: AccountType;
  name: Scalars['String'];
  age: Scalars['Int'];
  gender: AccountGender;
  avatar: Scalars['String'];
  introduction: Scalars['String'];
}>;

export type CreateAccountMutation = {
  __typename?: 'Mutation';
  createAccount: {
    __typename?: 'Account';
    id: string;
    email: string;
    password: string;
    type: AccountType;
    name: string;
    age: number;
    gender: AccountGender;
    avatar: string;
    introduction: string;
  };
};

export type UpdateAccountMutationVariables = Exact<{
  id: Scalars['ID'];
  email: Scalars['String'];
  password: Scalars['String'];
  type: AccountType;
  name: Scalars['String'];
  age: Scalars['Int'];
  gender: AccountGender;
  avatar: Scalars['String'];
  introduction: Scalars['String'];
}>;

export type UpdateAccountMutation = {
  __typename?: 'Mutation';
  updateAccount: {
    __typename?: 'Account';
    id: string;
    email: string;
    type: AccountType;
    name: string;
    age: number;
    gender: AccountGender;
    avatar: string;
    introduction: string;
  };
};

export type DeleteAccountMutationVariables = Exact<{
  id: Scalars['ID'];
}>;

export type DeleteAccountMutation = {
  __typename?: 'Mutation';
  deleteAccount: {
    __typename?: 'Account';
    id: string;
    email: string;
    type: AccountType;
    name: string;
    age: number;
    gender: AccountGender;
    avatar: string;
    introduction: string;
  };
};

export type LoginQueryMutationVariables = Exact<{
  email: Scalars['String'];
  password: Scalars['String'];
}>;

export type LoginQueryMutation = {
  __typename?: 'Mutation';
  loginQuery: {
    __typename?: 'Account';
    id: string;
    email: string;
    type: AccountType;
    name: string;
    age: number;
    gender: AccountGender;
    avatar: string;
    introduction: string;
  };
};

export type CreateCommentMutationVariables = Exact<{
  postId: Scalars['ID'];
  accountId: Scalars['ID'];
  body: Scalars['String'];
}>;

export type CreateCommentMutation = {
  __typename?: 'Mutation';
  createComment: {
    __typename?: 'Comment';
    id: string;
    postId: string;
    accountId: string;
    body: string;
  };
};

export type DeleteCommentMutationVariables = Exact<{
  id: Scalars['ID'];
}>;

export type DeleteCommentMutation = {
  __typename?: 'Mutation';
  deleteComment: {
    __typename?: 'Comment';
    id: string;
    postId: string;
    accountId: string;
    body: string;
  };
};

export type CreateFriendMutationVariables = Exact<{
  accountId: Scalars['ID'];
  friendId: Scalars['ID'];
}>;

export type CreateFriendMutation = {
  __typename?: 'Mutation';
  createFriend: { __typename?: 'Friend'; accountId: string; friendId: string };
};

export type DeleteFriendMutationVariables = Exact<{
  accountId: Scalars['ID'];
  friendId: Scalars['ID'];
}>;

export type DeleteFriendMutation = {
  __typename?: 'Mutation';
  deleteFriend: { __typename?: 'Friend'; accountId: string; friendId: string };
};

export type CreateLikeMutationVariables = Exact<{
  postId: Scalars['ID'];
  accountId: Scalars['ID'];
}>;

export type CreateLikeMutation = {
  __typename?: 'Mutation';
  createLike: { __typename?: 'Like'; postId: string; accountId: string };
};

export type DeleteLikeMutationVariables = Exact<{
  postId: Scalars['ID'];
  accountId: Scalars['ID'];
}>;

export type DeleteLikeMutation = {
  __typename?: 'Mutation';
  deleteLike: { __typename?: 'Like'; postId: string; accountId: string };
};

export type CreateMarkerMutationVariables = Exact<{
  postId: Scalars['ID'];
  title: Scalars['String'];
  lat: Scalars['String'];
  lng: Scalars['String'];
}>;

export type CreateMarkerMutation = {
  __typename?: 'Mutation';
  createMarker: {
    __typename?: 'Marker';
    id: string;
    postId: string;
    title: string;
    lat: string;
    lng: string;
  };
};

export type UpdateMarkerMutationVariables = Exact<{
  id: Scalars['ID'];
  title: Scalars['String'];
  lat: Scalars['String'];
  lng: Scalars['String'];
}>;

export type UpdateMarkerMutation = {
  __typename?: 'Mutation';
  updateMarker: {
    __typename?: 'Marker';
    id: string;
    postId: string;
    title: string;
    lat: string;
    lng: string;
  };
};

export type DeleteMarkerMutationVariables = Exact<{
  id: Scalars['ID'];
}>;

export type DeleteMarkerMutation = {
  __typename?: 'Mutation';
  deleteMarker: {
    __typename?: 'Marker';
    id: string;
    postId: string;
    title: string;
    lat: string;
    lng: string;
  };
};

export type CreateMuteMutationVariables = Exact<{
  accountId: Scalars['ID'];
  muteId: Scalars['ID'];
}>;

export type CreateMuteMutation = {
  __typename?: 'Mutation';
  createMute: { __typename?: 'Mute'; accountId: string; muteId: string };
};

export type DeleteMuteMutationVariables = Exact<{
  accountId: Scalars['ID'];
  muteId: Scalars['ID'];
}>;

export type DeleteMuteMutation = {
  __typename?: 'Mutation';
  deleteMute: { __typename?: 'Mute'; accountId: string; muteId: string };
};

export type CreatePostMutationVariables = Exact<{
  accountId: Scalars['ID'];
  title: Scalars['String'];
  body: Scalars['String'];
  img: Scalars['String'];
}>;

export type CreatePostMutation = {
  __typename?: 'Mutation';
  createPost: {
    __typename?: 'Post';
    id: string;
    accountId: string;
    title: string;
    body: string;
    img: string;
  };
};

export type UpdatePostMutationVariables = Exact<{
  id: Scalars['ID'];
  accountId: Scalars['ID'];
  title: Scalars['String'];
  body: Scalars['String'];
  img: Scalars['String'];
}>;

export type UpdatePostMutation = {
  __typename?: 'Mutation';
  updatePost: {
    __typename?: 'Post';
    id: string;
    accountId: string;
    title: string;
    body: string;
    img: string;
  };
};

export type DeletePostMutationVariables = Exact<{
  id: Scalars['ID'];
}>;

export type DeletePostMutation = {
  __typename?: 'Mutation';
  deletePost: {
    __typename?: 'Post';
    id: string;
    accountId: string;
    title: string;
    body: string;
    img: string;
  };
};

export type CreateRequestMutationVariables = Exact<{
  accountId: Scalars['ID'];
  targetAccountId: Scalars['ID'];
  status: RequestStatus;
}>;

export type CreateRequestMutation = {
  __typename?: 'Mutation';
  createRequest: {
    __typename?: 'Request';
    accountId: string;
    targetAccountId: string;
    status: RequestStatus;
  };
};

export type UpdateRequestMutationVariables = Exact<{
  accountId: Scalars['ID'];
  targetAccountId: Scalars['ID'];
  status: RequestStatus;
}>;

export type UpdateRequestMutation = {
  __typename?: 'Mutation';
  updateRequest: {
    __typename?: 'Request';
    accountId: string;
    targetAccountId: string;
    status: RequestStatus;
  };
};

export type DeleteRequestMutationVariables = Exact<{
  accountId: Scalars['ID'];
  targetAccountId: Scalars['ID'];
}>;

export type DeleteRequestMutation = {
  __typename?: 'Mutation';
  deleteRequest: {
    __typename?: 'Request';
    accountId: string;
    targetAccountId: string;
    status: RequestStatus;
  };
};

export type CreateSessionMutationVariables = Exact<{
  accountId: Scalars['ID'];
}>;

export type CreateSessionMutation = {
  __typename?: 'Mutation';
  createSession: { __typename?: 'Session'; accountId: string };
};

export type UpdateSessionMutationVariables = Exact<{
  accountId: Scalars['ID'];
}>;

export type UpdateSessionMutation = {
  __typename?: 'Mutation';
  updateSession: { __typename?: 'Session'; accountId: string; session: string };
};

export type DeleteSessionMutationVariables = Exact<{
  accountId: Scalars['ID'];
}>;

export type DeleteSessionMutation = {
  __typename?: 'Mutation';
  deleteSession: { __typename?: 'Session'; accountId: string };
};

export type AccountPageInfoQueryVariables = Exact<{
  id: Scalars['ID'];
}>;

export type AccountPageInfoQuery = {
  __typename?: 'Query';
  accountPageInfo: {
    __typename?: 'Account';
    id: string;
    name: string;
    age: number;
    gender: AccountGender;
    avatar: string;
    introduction: string;
    post: Array<{
      __typename?: 'Post';
      id: string;
      title: string;
      updatedAt: any;
      like: Array<{ __typename?: 'Like'; accountId: string } | null>;
    } | null>;
    friend: Array<{
      __typename?: 'Friend';
      friend: {
        __typename?: 'Account';
        id: string;
        name: string;
        avatar: string;
      };
    } | null>;
  };
};

export type MyPageInfoQueryVariables = Exact<{
  id: Scalars['ID'];
}>;

export type MyPageInfoQuery = {
  __typename?: 'Query';
  myPageInfo: {
    __typename?: 'Account';
    id: string;
    email: string;
    password: string;
    type: AccountType;
    name: string;
    age: number;
    gender: AccountGender;
    avatar: string;
    introduction: string;
    post: Array<{
      __typename?: 'Post';
      id: string;
      title: string;
      createdAt: any;
      updatedAt: any;
      like: Array<{ __typename?: 'Like'; accountId: string } | null>;
    } | null>;
    like: Array<{
      __typename?: 'Like';
      post: {
        __typename?: 'Post';
        id: string;
        accountId: string;
        title: string;
        updatedAt: any;
      };
      account: {
        __typename?: 'Account';
        id: string;
        name: string;
        avatar: string;
      };
    } | null>;
    friend: Array<{
      __typename?: 'Friend';
      friend: {
        __typename?: 'Account';
        id: string;
        name: string;
        avatar: string;
      };
    } | null>;
    mute: Array<{
      __typename?: 'Mute';
      mute: {
        __typename?: 'Account';
        id: string;
        name: string;
        avatar: string;
      };
    } | null>;
  };
};

export type LikesByPostQueryVariables = Exact<{
  postId: Scalars['ID'];
}>;

export type LikesByPostQuery = {
  __typename?: 'Query';
  likesByPost: Array<{
    __typename?: 'Like';
    createdAt: any;
    account: {
      __typename?: 'Account';
      id: string;
      name: string;
      avatar: string;
    };
  }>;
};

export type AllMarkersQueryVariables = Exact<{ [key: string]: never }>;

export type AllMarkersQuery = {
  __typename?: 'Query';
  allMarkers: Array<{
    __typename?: 'Marker';
    id: string;
    title: string;
    lat: string;
    lng: string;
    post: {
      __typename?: 'Post';
      id: string;
      title: string;
      body: string;
      img: string;
      account: {
        __typename?: 'Account';
        id: string;
        name: string;
        avatar: string;
      };
      like: Array<{ __typename?: 'Like'; accountId: string } | null>;
    };
  }>;
};

export type AllPostsQueryVariables = Exact<{ [key: string]: never }>;

export type AllPostsQuery = {
  __typename?: 'Query';
  allPosts: Array<{
    __typename?: 'Post';
    id: string;
    title: string;
    body: string;
    img: string;
    createdAt: any;
    updatedAt: any;
    account: {
      __typename?: 'Account';
      id: string;
      name: string;
      avatar: string;
    };
    marker?: {
      __typename?: 'Marker';
      id: string;
      title: string;
      lat: string;
      lng: string;
    } | null;
    like: Array<{ __typename?: 'Like'; accountId: string } | null>;
    comment: Array<{
      __typename?: 'Comment';
      id: string;
      body: string;
      account: {
        __typename?: 'Account';
        id: string;
        name: string;
        avatar: string;
      };
    } | null>;
  }>;
};

export type RequestsByAccountIdQueryVariables = Exact<{
  accountId: Scalars['ID'];
}>;

export type RequestsByAccountIdQuery = {
  __typename?: 'Query';
  requestsByAccountId: Array<{
    __typename?: 'Request';
    targetAccount: {
      __typename?: 'Account';
      id: string;
      name: string;
      avatar: string;
    };
  }>;
};

export type RequestsByTargetIdQueryVariables = Exact<{
  targetAccountId: Scalars['ID'];
}>;

export type RequestsByTargetIdQuery = {
  __typename?: 'Query';
  requestsByTargetId: Array<{
    __typename?: 'Request';
    account: {
      __typename?: 'Account';
      id: string;
      name: string;
      avatar: string;
    };
  }>;
};

export type SessionByAccountIdQueryVariables = Exact<{
  accountId: Scalars['ID'];
}>;

export type SessionByAccountIdQuery = {
  __typename?: 'Query';
  sessionByAccountId: { __typename?: 'Session'; session: string };
};

import { IntrospectionQuery } from 'graphql';
export default {
  __schema: {
    queryType: {
      name: 'Query',
    },
    mutationType: {
      name: 'Mutation',
    },
    subscriptionType: null,
    types: [
      {
        kind: 'OBJECT',
        name: 'Account',
        fields: [
          {
            name: 'age',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'avatar',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'comment',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'OBJECT',
                  name: 'Comment',
                  ofType: null,
                },
              },
            },
            args: [],
          },
          {
            name: 'createdAt',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'email',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'friend',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'OBJECT',
                  name: 'Friend',
                  ofType: null,
                },
              },
            },
            args: [],
          },
          {
            name: 'gender',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'id',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'introduction',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'like',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'OBJECT',
                  name: 'Like',
                  ofType: null,
                },
              },
            },
            args: [],
          },
          {
            name: 'mute',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'OBJECT',
                  name: 'Mute',
                  ofType: null,
                },
              },
            },
            args: [],
          },
          {
            name: 'name',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'password',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'post',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'OBJECT',
                  name: 'Post',
                  ofType: null,
                },
              },
            },
            args: [],
          },
          {
            name: 'type',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'updatedAt',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'Comment',
        fields: [
          {
            name: 'account',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: 'accountId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'body',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'createdAt',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'id',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'post',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Post',
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: 'postId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'updatedAt',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'File',
        fields: [
          {
            name: 'path',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'Friend',
        fields: [
          {
            name: 'accountId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'friend',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: 'friendId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'Like',
        fields: [
          {
            name: 'account',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: 'accountId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'createdAt',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'post',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Post',
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: 'postId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'Marker',
        fields: [
          {
            name: 'createdAt',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'id',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'lat',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'lng',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'post',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Post',
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: 'postId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'title',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'updatedAt',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'Mutation',
        fields: [
          {
            name: 'createAccount',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'createComment',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Comment',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'createFriend',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Friend',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'createLike',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Like',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'createMarker',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Marker',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'createMute',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Mute',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'createPost',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Post',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'createRequest',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Request',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'createSession',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Session',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'deleteAccount',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [
              {
                name: 'id',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'deleteComment',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Comment',
                ofType: null,
              },
            },
            args: [
              {
                name: 'id',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'deleteFriend',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Friend',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'deleteLike',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Like',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'deleteMarker',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Marker',
                ofType: null,
              },
            },
            args: [
              {
                name: 'id',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'deleteMute',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Mute',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'SCALAR',
                  name: 'Any',
                },
              },
            ],
          },
          {
            name: 'deletePost',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Post',
                ofType: null,
              },
            },
            args: [
              {
                name: 'id',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'deleteRequest',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Request',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'deleteSession',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Session',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'loginQuery',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'updateAccount',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [
              {
                name: 'id',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'updateMarker',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Marker',
                ofType: null,
              },
            },
            args: [
              {
                name: 'id',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'updatePost',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Post',
                ofType: null,
              },
            },
            args: [
              {
                name: 'id',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'updateRequest',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Request',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'updateSession',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Session',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'uploadFile',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'File',
                ofType: null,
              },
            },
            args: [
              {
                name: 'input',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'Mute',
        fields: [
          {
            name: 'accountId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'mute',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: 'muteId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'Post',
        fields: [
          {
            name: 'account',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: 'accountId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'body',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'comment',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'OBJECT',
                  name: 'Comment',
                  ofType: null,
                },
              },
            },
            args: [],
          },
          {
            name: 'createdAt',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'id',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'img',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'like',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'OBJECT',
                  name: 'Like',
                  ofType: null,
                },
              },
            },
            args: [],
          },
          {
            name: 'marker',
            type: {
              kind: 'OBJECT',
              name: 'Marker',
              ofType: null,
            },
            args: [],
          },
          {
            name: 'title',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'updatedAt',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'Query',
        fields: [
          {
            name: 'accountPageInfo',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [
              {
                name: 'id',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'allMarkers',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'OBJECT',
                    name: 'Marker',
                    ofType: null,
                  },
                },
              },
            },
            args: [],
          },
          {
            name: 'allPosts',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'OBJECT',
                    name: 'Post',
                    ofType: null,
                  },
                },
              },
            },
            args: [],
          },
          {
            name: 'likesByPost',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'OBJECT',
                    name: 'Like',
                    ofType: null,
                  },
                },
              },
            },
            args: [
              {
                name: 'postId',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'myPageInfo',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [
              {
                name: 'id',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'requestsByAccountId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'OBJECT',
                    name: 'Request',
                    ofType: null,
                  },
                },
              },
            },
            args: [
              {
                name: 'accountId',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'requestsByTargetId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'LIST',
                ofType: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'OBJECT',
                    name: 'Request',
                    ofType: null,
                  },
                },
              },
            },
            args: [
              {
                name: 'targetAccountId',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
          {
            name: 'sessionByAccountId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Session',
                ofType: null,
              },
            },
            args: [
              {
                name: 'accountId',
                type: {
                  kind: 'NON_NULL',
                  ofType: {
                    kind: 'SCALAR',
                    name: 'Any',
                  },
                },
              },
            ],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'Request',
        fields: [
          {
            name: 'account',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: 'accountId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'status',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'targetAccount',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'OBJECT',
                name: 'Account',
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: 'targetAccountId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'OBJECT',
        name: 'Session',
        fields: [
          {
            name: 'accountId',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'session',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
          {
            name: 'updatedAt',
            type: {
              kind: 'NON_NULL',
              ofType: {
                kind: 'SCALAR',
                name: 'Any',
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: 'SCALAR',
        name: 'Any',
      },
    ],
    directives: [],
  },
} as unknown as IntrospectionQuery;

export const CreateAccountDocument = gql`
  mutation createAccount(
    $email: String!
    $password: String!
    $type: AccountType!
    $name: String!
    $age: Int!
    $gender: AccountGender!
    $avatar: String!
    $introduction: String!
  ) {
    createAccount(
      input: {
        email: $email
        password: $password
        type: $type
        name: $name
        age: $age
        gender: $gender
        avatar: $avatar
        introduction: $introduction
      }
    ) {
      id
      email
      password
      type
      name
      age
      gender
      avatar
      introduction
    }
  }
`;

export function useCreateAccountMutation() {
  return Urql.useMutation<
    CreateAccountMutation,
    CreateAccountMutationVariables
  >(CreateAccountDocument);
}
export const UpdateAccountDocument = gql`
  mutation updateAccount(
    $id: ID!
    $email: String!
    $password: String!
    $type: AccountType!
    $name: String!
    $age: Int!
    $gender: AccountGender!
    $avatar: String!
    $introduction: String!
  ) {
    updateAccount(
      id: $id
      input: {
        email: $email
        password: $password
        type: $type
        name: $name
        age: $age
        gender: $gender
        avatar: $avatar
        introduction: $introduction
      }
    ) {
      id
      email
      type
      name
      age
      gender
      avatar
      introduction
    }
  }
`;

export function useUpdateAccountMutation() {
  return Urql.useMutation<
    UpdateAccountMutation,
    UpdateAccountMutationVariables
  >(UpdateAccountDocument);
}
export const DeleteAccountDocument = gql`
  mutation deleteAccount($id: ID!) {
    deleteAccount(id: $id) {
      id
      email
      type
      name
      age
      gender
      avatar
      introduction
    }
  }
`;

export function useDeleteAccountMutation() {
  return Urql.useMutation<
    DeleteAccountMutation,
    DeleteAccountMutationVariables
  >(DeleteAccountDocument);
}
export const LoginQueryDocument = gql`
  mutation loginQuery($email: String!, $password: String!) {
    loginQuery(input: { email: $email, password: $password }) {
      id
      email
      type
      name
      age
      gender
      avatar
      introduction
    }
  }
`;

export function useLoginQueryMutation() {
  return Urql.useMutation<LoginQueryMutation, LoginQueryMutationVariables>(
    LoginQueryDocument,
  );
}
export const CreateCommentDocument = gql`
  mutation createComment($postId: ID!, $accountId: ID!, $body: String!) {
    createComment(
      input: { postId: $postId, accountId: $accountId, body: $body }
    ) {
      id
      postId
      accountId
      body
    }
  }
`;

export function useCreateCommentMutation() {
  return Urql.useMutation<
    CreateCommentMutation,
    CreateCommentMutationVariables
  >(CreateCommentDocument);
}
export const DeleteCommentDocument = gql`
  mutation deleteComment($id: ID!) {
    deleteComment(id: $id) {
      id
      postId
      accountId
      body
    }
  }
`;

export function useDeleteCommentMutation() {
  return Urql.useMutation<
    DeleteCommentMutation,
    DeleteCommentMutationVariables
  >(DeleteCommentDocument);
}
export const CreateFriendDocument = gql`
  mutation createFriend($accountId: ID!, $friendId: ID!) {
    createFriend(input: { accountId: $accountId, friendId: $friendId }) {
      accountId
      friendId
    }
  }
`;

export function useCreateFriendMutation() {
  return Urql.useMutation<CreateFriendMutation, CreateFriendMutationVariables>(
    CreateFriendDocument,
  );
}
export const DeleteFriendDocument = gql`
  mutation deleteFriend($accountId: ID!, $friendId: ID!) {
    deleteFriend(input: { accountId: $accountId, friendId: $friendId }) {
      accountId
      friendId
    }
  }
`;

export function useDeleteFriendMutation() {
  return Urql.useMutation<DeleteFriendMutation, DeleteFriendMutationVariables>(
    DeleteFriendDocument,
  );
}
export const CreateLikeDocument = gql`
  mutation createLike($postId: ID!, $accountId: ID!) {
    createLike(input: { postId: $postId, accountId: $accountId }) {
      postId
      accountId
    }
  }
`;

export function useCreateLikeMutation() {
  return Urql.useMutation<CreateLikeMutation, CreateLikeMutationVariables>(
    CreateLikeDocument,
  );
}
export const DeleteLikeDocument = gql`
  mutation deleteLike($postId: ID!, $accountId: ID!) {
    deleteLike(input: { postId: $postId, accountId: $accountId }) {
      postId
      accountId
    }
  }
`;

export function useDeleteLikeMutation() {
  return Urql.useMutation<DeleteLikeMutation, DeleteLikeMutationVariables>(
    DeleteLikeDocument,
  );
}
export const CreateMarkerDocument = gql`
  mutation createMarker(
    $postId: ID!
    $title: String!
    $lat: String!
    $lng: String!
  ) {
    createMarker(
      input: { postId: $postId, title: $title, lat: $lat, lng: $lng }
    ) {
      id
      postId
      title
      lat
      lng
    }
  }
`;

export function useCreateMarkerMutation() {
  return Urql.useMutation<CreateMarkerMutation, CreateMarkerMutationVariables>(
    CreateMarkerDocument,
  );
}
export const UpdateMarkerDocument = gql`
  mutation updateMarker(
    $id: ID!
    $title: String!
    $lat: String!
    $lng: String!
  ) {
    updateMarker(id: $id, input: { title: $title, lat: $lat, lng: $lng }) {
      id
      postId
      title
      lat
      lng
    }
  }
`;

export function useUpdateMarkerMutation() {
  return Urql.useMutation<UpdateMarkerMutation, UpdateMarkerMutationVariables>(
    UpdateMarkerDocument,
  );
}
export const DeleteMarkerDocument = gql`
  mutation deleteMarker($id: ID!) {
    deleteMarker(id: $id) {
      id
      postId
      title
      lat
      lng
    }
  }
`;

export function useDeleteMarkerMutation() {
  return Urql.useMutation<DeleteMarkerMutation, DeleteMarkerMutationVariables>(
    DeleteMarkerDocument,
  );
}
export const CreateMuteDocument = gql`
  mutation createMute($accountId: ID!, $muteId: ID!) {
    createMute(input: { accountId: $accountId, muteId: $muteId }) {
      accountId
      muteId
    }
  }
`;

export function useCreateMuteMutation() {
  return Urql.useMutation<CreateMuteMutation, CreateMuteMutationVariables>(
    CreateMuteDocument,
  );
}
export const DeleteMuteDocument = gql`
  mutation deleteMute($accountId: ID!, $muteId: ID!) {
    deleteMute(input: { accountId: $accountId, muteId: $muteId }) {
      accountId
      muteId
    }
  }
`;

export function useDeleteMuteMutation() {
  return Urql.useMutation<DeleteMuteMutation, DeleteMuteMutationVariables>(
    DeleteMuteDocument,
  );
}
export const CreatePostDocument = gql`
  mutation createPost(
    $accountId: ID!
    $title: String!
    $body: String!
    $img: String!
  ) {
    createPost(
      input: { accountId: $accountId, title: $title, body: $body, img: $img }
    ) {
      id
      accountId
      title
      body
      img
    }
  }
`;

export function useCreatePostMutation() {
  return Urql.useMutation<CreatePostMutation, CreatePostMutationVariables>(
    CreatePostDocument,
  );
}
export const UpdatePostDocument = gql`
  mutation updatePost(
    $id: ID!
    $accountId: ID!
    $title: String!
    $body: String!
    $img: String!
  ) {
    updatePost(
      id: $id
      input: { accountId: $accountId, title: $title, body: $body, img: $img }
    ) {
      id
      accountId
      title
      body
      img
    }
  }
`;

export function useUpdatePostMutation() {
  return Urql.useMutation<UpdatePostMutation, UpdatePostMutationVariables>(
    UpdatePostDocument,
  );
}
export const DeletePostDocument = gql`
  mutation deletePost($id: ID!) {
    deletePost(id: $id) {
      id
      accountId
      title
      body
      img
    }
  }
`;

export function useDeletePostMutation() {
  return Urql.useMutation<DeletePostMutation, DeletePostMutationVariables>(
    DeletePostDocument,
  );
}
export const CreateRequestDocument = gql`
  mutation createRequest(
    $accountId: ID!
    $targetAccountId: ID!
    $status: RequestStatus!
  ) {
    createRequest(
      input: {
        accountId: $accountId
        targetAccountId: $targetAccountId
        status: $status
      }
    ) {
      accountId
      targetAccountId
      status
    }
  }
`;

export function useCreateRequestMutation() {
  return Urql.useMutation<
    CreateRequestMutation,
    CreateRequestMutationVariables
  >(CreateRequestDocument);
}
export const UpdateRequestDocument = gql`
  mutation updateRequest(
    $accountId: ID!
    $targetAccountId: ID!
    $status: RequestStatus!
  ) {
    updateRequest(
      input: {
        accountId: $accountId
        targetAccountId: $targetAccountId
        status: $status
      }
    ) {
      accountId
      targetAccountId
      status
    }
  }
`;

export function useUpdateRequestMutation() {
  return Urql.useMutation<
    UpdateRequestMutation,
    UpdateRequestMutationVariables
  >(UpdateRequestDocument);
}
export const DeleteRequestDocument = gql`
  mutation deleteRequest($accountId: ID!, $targetAccountId: ID!) {
    deleteRequest(
      input: { accountId: $accountId, targetAccountId: $targetAccountId }
    ) {
      accountId
      targetAccountId
      status
    }
  }
`;

export function useDeleteRequestMutation() {
  return Urql.useMutation<
    DeleteRequestMutation,
    DeleteRequestMutationVariables
  >(DeleteRequestDocument);
}
export const CreateSessionDocument = gql`
  mutation createSession($accountId: ID!) {
    createSession(input: { accountId: $accountId }) {
      accountId
    }
  }
`;

export function useCreateSessionMutation() {
  return Urql.useMutation<
    CreateSessionMutation,
    CreateSessionMutationVariables
  >(CreateSessionDocument);
}
export const UpdateSessionDocument = gql`
  mutation updateSession($accountId: ID!) {
    updateSession(input: { accountId: $accountId }) {
      accountId
      session
    }
  }
`;

export function useUpdateSessionMutation() {
  return Urql.useMutation<
    UpdateSessionMutation,
    UpdateSessionMutationVariables
  >(UpdateSessionDocument);
}
export const DeleteSessionDocument = gql`
  mutation deleteSession($accountId: ID!) {
    deleteSession(input: { accountId: $accountId }) {
      accountId
    }
  }
`;

export function useDeleteSessionMutation() {
  return Urql.useMutation<
    DeleteSessionMutation,
    DeleteSessionMutationVariables
  >(DeleteSessionDocument);
}
export const AccountPageInfoDocument = gql`
  query accountPageInfo($id: ID!) {
    accountPageInfo(id: $id) {
      id
      name
      age
      gender
      avatar
      introduction
      post {
        id
        title
        updatedAt
        like {
          accountId
        }
      }
      friend {
        friend {
          id
          name
          avatar
        }
      }
    }
  }
`;

export function useAccountPageInfoQuery(
  options: Omit<Urql.UseQueryArgs<AccountPageInfoQueryVariables>, 'query'>,
) {
  return Urql.useQuery<AccountPageInfoQuery>({
    query: AccountPageInfoDocument,
    ...options,
  });
}
export const MyPageInfoDocument = gql`
  query myPageInfo($id: ID!) {
    myPageInfo(id: $id) {
      id
      email
      password
      type
      name
      age
      gender
      avatar
      introduction
      post {
        id
        title
        createdAt
        updatedAt
        like {
          accountId
        }
      }
      like {
        post {
          id
          accountId
          title
          updatedAt
        }
        account {
          id
          name
          avatar
        }
      }
      friend {
        friend {
          id
          name
          avatar
        }
      }
      mute {
        mute {
          id
          name
          avatar
        }
      }
    }
  }
`;

export function useMyPageInfoQuery(
  options: Omit<Urql.UseQueryArgs<MyPageInfoQueryVariables>, 'query'>,
) {
  return Urql.useQuery<MyPageInfoQuery>({
    query: MyPageInfoDocument,
    ...options,
  });
}
export const LikesByPostDocument = gql`
  query likesByPost($postId: ID!) {
    likesByPost(postId: $postId) {
      account {
        id
        name
        avatar
      }
      createdAt
    }
  }
`;

export function useLikesByPostQuery(
  options: Omit<Urql.UseQueryArgs<LikesByPostQueryVariables>, 'query'>,
) {
  return Urql.useQuery<LikesByPostQuery>({
    query: LikesByPostDocument,
    ...options,
  });
}
export const AllMarkersDocument = gql`
  query allMarkers {
    allMarkers {
      id
      title
      lat
      lng
      post {
        id
        title
        body
        img
        account {
          id
          name
          avatar
        }
        like {
          accountId
        }
      }
    }
  }
`;

export function useAllMarkersQuery(
  options?: Omit<Urql.UseQueryArgs<AllMarkersQueryVariables>, 'query'>,
) {
  return Urql.useQuery<AllMarkersQuery>({
    query: AllMarkersDocument,
    ...options,
  });
}
export const AllPostsDocument = gql`
  query allPosts {
    allPosts {
      id
      title
      body
      img
      createdAt
      updatedAt
      account {
        id
        name
        avatar
      }
      marker {
        id
        title
        lat
        lng
      }
      like {
        accountId
      }
      comment {
        id
        body
        account {
          id
          name
          avatar
        }
      }
    }
  }
`;

export function useAllPostsQuery(
  options?: Omit<Urql.UseQueryArgs<AllPostsQueryVariables>, 'query'>,
) {
  return Urql.useQuery<AllPostsQuery>({ query: AllPostsDocument, ...options });
}
export const RequestsByAccountIdDocument = gql`
  query requestsByAccountId($accountId: ID!) {
    requestsByAccountId(accountId: $accountId) {
      targetAccount {
        id
        name
        avatar
      }
    }
  }
`;

export function useRequestsByAccountIdQuery(
  options: Omit<Urql.UseQueryArgs<RequestsByAccountIdQueryVariables>, 'query'>,
) {
  return Urql.useQuery<RequestsByAccountIdQuery>({
    query: RequestsByAccountIdDocument,
    ...options,
  });
}
export const RequestsByTargetIdDocument = gql`
  query requestsByTargetId($targetAccountId: ID!) {
    requestsByTargetId(targetAccountId: $targetAccountId) {
      account {
        id
        name
        avatar
      }
    }
  }
`;

export function useRequestsByTargetIdQuery(
  options: Omit<Urql.UseQueryArgs<RequestsByTargetIdQueryVariables>, 'query'>,
) {
  return Urql.useQuery<RequestsByTargetIdQuery>({
    query: RequestsByTargetIdDocument,
    ...options,
  });
}
export const SessionByAccountIdDocument = gql`
  query sessionByAccountId($accountId: ID!) {
    sessionByAccountId(accountId: $accountId) {
      session
    }
  }
`;

export function useSessionByAccountIdQuery(
  options: Omit<Urql.UseQueryArgs<SessionByAccountIdQueryVariables>, 'query'>,
) {
  return Urql.useQuery<SessionByAccountIdQuery>({
    query: SessionByAccountIdDocument,
    ...options,
  });
}
