import { IntrospectionQuery } from "graphql";

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
  __typename?: "Account";
  age: Scalars["Int"];
  avatar: Scalars["String"];
  comment?: Maybe<Array<Maybe<Comment>>>;
  createdAt: Scalars["Time"];
  email: Scalars["String"];
  friend?: Maybe<Array<Maybe<Friend>>>;
  gender: AccountGender;
  id: Scalars["ID"];
  introduction: Scalars["String"];
  like?: Maybe<Array<Maybe<Like>>>;
  mute?: Maybe<Array<Maybe<Mute>>>;
  name: Scalars["String"];
  password: Scalars["String"];
  post?: Maybe<Array<Maybe<Post>>>;
  type: AccountType;
  updatedAt: Scalars["Time"];
};

export enum AccountGender {
  Female = "female",
  Male = "male",
  None = "none",
}

export type AccountInput = {
  age: Scalars["Int"];
  avatar: Scalars["String"];
  email: Scalars["String"];
  gender: AccountGender;
  introduction: Scalars["String"];
  name: Scalars["String"];
  password: Scalars["String"];
  type: AccountType;
};

export enum AccountType {
  Active = "active",
  Admin = "admin",
  Inactive = "inactive",
}

export type Comment = {
  __typename?: "Comment";
  account: Account;
  account_id: Scalars["ID"];
  body: Scalars["String"];
  createdAt: Scalars["Time"];
  id: Scalars["ID"];
  post: Post;
  post_id: Scalars["ID"];
  updatedAt: Scalars["Time"];
};

export type CommentInput = {
  account_id: Scalars["ID"];
  body: Scalars["String"];
  post_id: Scalars["ID"];
};

export type File = {
  __typename?: "File";
  path: Scalars["String"];
};

export type Friend = {
  __typename?: "Friend";
  account_id: Scalars["ID"];
  friend: Account;
  friend_id: Scalars["ID"];
};

export type FriendInput = {
  account_id: Scalars["ID"];
  friend_id: Scalars["ID"];
};

export type Like = {
  __typename?: "Like";
  account: Account;
  account_id: Scalars["ID"];
  post: Post;
  post_id: Scalars["ID"];
};

export type LikeInput = {
  account_id: Scalars["ID"];
  post_id: Scalars["ID"];
};

export type LoginInput = {
  email: Scalars["String"];
  password: Scalars["String"];
};

export type Marker = {
  __typename?: "Marker";
  createdAt: Scalars["Time"];
  id: Scalars["ID"];
  lat: Scalars["String"];
  lng: Scalars["String"];
  post: Post;
  post_id: Scalars["ID"];
  title: Scalars["String"];
  updatedAt: Scalars["Time"];
};

export type MarkerInput = {
  lat: Scalars["String"];
  lng: Scalars["String"];
  post_id: Scalars["ID"];
  title: Scalars["String"];
};

export type Mutation = {
  __typename?: "Mutation";
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
  login: Account;
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
  id: Scalars["ID"];
};

export type MutationDeleteCommentArgs = {
  id: Scalars["ID"];
};

export type MutationDeleteFriendArgs = {
  input: FriendInput;
};

export type MutationDeleteLikeArgs = {
  input: LikeInput;
};

export type MutationDeleteMarkerArgs = {
  id: Scalars["ID"];
};

export type MutationDeleteMuteArgs = {
  input?: InputMaybe<MuteInput>;
};

export type MutationDeletePostArgs = {
  id: Scalars["ID"];
};

export type MutationDeleteRequestArgs = {
  input: RequestInput;
};

export type MutationDeleteSessionArgs = {
  account_id: Scalars["ID"];
};

export type MutationLoginArgs = {
  input: LoginInput;
};

export type MutationUpdateAccountArgs = {
  id: Scalars["ID"];
  input: AccountInput;
};

export type MutationUpdateMarkerArgs = {
  id: Scalars["ID"];
  input: MarkerInput;
};

export type MutationUpdatePostArgs = {
  id: Scalars["ID"];
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
  __typename?: "Mute";
  account_id: Scalars["ID"];
  mute: Account;
  mute_id: Scalars["ID"];
};

export type MuteInput = {
  account_id: Scalars["ID"];
  mute_id: Scalars["ID"];
};

export type Post = {
  __typename?: "Post";
  account: Account;
  account_id: Scalars["ID"];
  body: Scalars["String"];
  comment?: Maybe<Array<Maybe<Comment>>>;
  createdAt: Scalars["Time"];
  id: Scalars["ID"];
  img: Scalars["String"];
  like?: Maybe<Array<Maybe<Like>>>;
  marker?: Maybe<Marker>;
  title: Scalars["String"];
  updatedAt: Scalars["Time"];
};

export type PostInput = {
  account_id: Scalars["ID"];
  body: Scalars["String"];
  img: Scalars["String"];
  title: Scalars["String"];
};

export type Query = {
  __typename?: "Query";
  getAccountByID: Account;
  getAllAccounts: Array<Maybe<Account>>;
  getAllMarkers: Array<Maybe<Marker>>;
  getAllPosts: Array<Maybe<Post>>;
  getCommentByAccountID: Array<Maybe<Comment>>;
  getFriendsByAccountID: Array<Maybe<Friend>>;
  getLikesByAccountID: Array<Maybe<Like>>;
  getMutesByAccountID: Array<Maybe<Mute>>;
  getPostsByAccountID: Array<Maybe<Post>>;
  getRequestsByAccountID: Array<Maybe<Request>>;
  getRequestsByTargetID: Array<Maybe<Request>>;
  getSessionByAccountID: Session;
};

export type QueryGetAccountByIdArgs = {
  id: Scalars["ID"];
};

export type QueryGetCommentByAccountIdArgs = {
  account_id: Scalars["ID"];
};

export type QueryGetFriendsByAccountIdArgs = {
  account_id: Scalars["ID"];
};

export type QueryGetLikesByAccountIdArgs = {
  account_id: Scalars["ID"];
};

export type QueryGetMutesByAccountIdArgs = {
  account_id: Scalars["ID"];
};

export type QueryGetPostsByAccountIdArgs = {
  account_id: Scalars["ID"];
};

export type QueryGetRequestsByAccountIdArgs = {
  account_id: Scalars["ID"];
};

export type QueryGetRequestsByTargetIdArgs = {
  target_account_id: Scalars["ID"];
};

export type QueryGetSessionByAccountIdArgs = {
  account_id: Scalars["ID"];
};

export type Request = {
  __typename?: "Request";
  account: Account;
  account_id: Scalars["ID"];
  status: RequestStatus;
  target_account: Account;
  target_account_id: Scalars["ID"];
};

export type RequestInput = {
  account_id: Scalars["ID"];
  status: RequestStatus;
  target_account_id: Scalars["ID"];
};

export enum RequestStatus {
  Accept = "accept",
  BreakAccept = "break_accept",
  BreakDeny = "break_deny",
  BreakInProcess = "break_in_process",
  Cancel = "cancel",
  Deny = "deny",
  InProcess = "in_process",
}

export type Session = {
  __typename?: "Session";
  account_id: Scalars["ID"];
  session: Scalars["String"];
  updatedAt: Scalars["Time"];
};

export type SessionInput = {
  account_id: Scalars["ID"];
  session: Scalars["String"];
};

export type UploadFile = {
  content: Scalars["Upload"];
};

export default {
  __schema: {
    queryType: {
      name: "Query",
    },
    mutationType: {
      name: "Mutation",
    },
    subscriptionType: null,
    types: [
      {
        kind: "OBJECT",
        name: "Account",
        fields: [
          {
            name: "age",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "avatar",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "comment",
            type: {
              kind: "LIST",
              ofType: {
                kind: "OBJECT",
                name: "Comment",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "createdAt",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "email",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "friend",
            type: {
              kind: "LIST",
              ofType: {
                kind: "OBJECT",
                name: "Friend",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "gender",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "introduction",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "like",
            type: {
              kind: "LIST",
              ofType: {
                kind: "OBJECT",
                name: "Like",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "mute",
            type: {
              kind: "LIST",
              ofType: {
                kind: "OBJECT",
                name: "Mute",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "name",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "password",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "post",
            type: {
              kind: "LIST",
              ofType: {
                kind: "OBJECT",
                name: "Post",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "type",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "updatedAt",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "Comment",
        fields: [
          {
            name: "account",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "account_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "body",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "createdAt",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "post",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Post",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "post_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "updatedAt",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "File",
        fields: [
          {
            name: "path",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "Friend",
        fields: [
          {
            name: "account_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "friend",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "friend_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "Like",
        fields: [
          {
            name: "account",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "account_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "post",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Post",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "post_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "Marker",
        fields: [
          {
            name: "createdAt",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "lat",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "lng",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "post",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Post",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "post_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "title",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "updatedAt",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "Mutation",
        fields: [
          {
            name: "createAccount",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "createComment",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Comment",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "createFriend",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Friend",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "createLike",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Like",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "createMarker",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Marker",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "createMute",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Mute",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "createPost",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Post",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "createRequest",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Request",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "createSession",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Session",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "deleteAccount",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [
              {
                name: "id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "deleteComment",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Comment",
                ofType: null,
              },
            },
            args: [
              {
                name: "id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "deleteFriend",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Friend",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "deleteLike",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Like",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "deleteMarker",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Marker",
                ofType: null,
              },
            },
            args: [
              {
                name: "id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "deleteMute",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Mute",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "SCALAR",
                  name: "Any",
                },
              },
            ],
          },
          {
            name: "deletePost",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Post",
                ofType: null,
              },
            },
            args: [
              {
                name: "id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "deleteRequest",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Request",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "deleteSession",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Session",
                ofType: null,
              },
            },
            args: [
              {
                name: "account_id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "login",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "updateAccount",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [
              {
                name: "id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "updateMarker",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Marker",
                ofType: null,
              },
            },
            args: [
              {
                name: "id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "updatePost",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Post",
                ofType: null,
              },
            },
            args: [
              {
                name: "id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "updateRequest",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Request",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "updateSession",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Session",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "uploadFile",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "File",
                ofType: null,
              },
            },
            args: [
              {
                name: "input",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "Mute",
        fields: [
          {
            name: "account_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "mute",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "mute_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "Post",
        fields: [
          {
            name: "account",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "account_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "body",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "comment",
            type: {
              kind: "LIST",
              ofType: {
                kind: "OBJECT",
                name: "Comment",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "createdAt",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "img",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "like",
            type: {
              kind: "LIST",
              ofType: {
                kind: "OBJECT",
                name: "Like",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "marker",
            type: {
              kind: "OBJECT",
              name: "Marker",
              ofType: null,
            },
            args: [],
          },
          {
            name: "title",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "updatedAt",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "Query",
        fields: [
          {
            name: "getAccountByID",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [
              {
                name: "id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "getAllAccounts",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "LIST",
                ofType: {
                  kind: "OBJECT",
                  name: "Account",
                  ofType: null,
                },
              },
            },
            args: [],
          },
          {
            name: "getAllMarkers",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "LIST",
                ofType: {
                  kind: "OBJECT",
                  name: "Marker",
                  ofType: null,
                },
              },
            },
            args: [],
          },
          {
            name: "getAllPosts",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "LIST",
                ofType: {
                  kind: "OBJECT",
                  name: "Post",
                  ofType: null,
                },
              },
            },
            args: [],
          },
          {
            name: "getCommentByAccountID",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "LIST",
                ofType: {
                  kind: "OBJECT",
                  name: "Comment",
                  ofType: null,
                },
              },
            },
            args: [
              {
                name: "account_id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "getFriendsByAccountID",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "LIST",
                ofType: {
                  kind: "OBJECT",
                  name: "Friend",
                  ofType: null,
                },
              },
            },
            args: [
              {
                name: "account_id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "getLikesByAccountID",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "LIST",
                ofType: {
                  kind: "OBJECT",
                  name: "Like",
                  ofType: null,
                },
              },
            },
            args: [
              {
                name: "account_id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "getMutesByAccountID",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "LIST",
                ofType: {
                  kind: "OBJECT",
                  name: "Mute",
                  ofType: null,
                },
              },
            },
            args: [
              {
                name: "account_id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "getPostsByAccountID",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "LIST",
                ofType: {
                  kind: "OBJECT",
                  name: "Post",
                  ofType: null,
                },
              },
            },
            args: [
              {
                name: "account_id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "getRequestsByAccountID",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "LIST",
                ofType: {
                  kind: "OBJECT",
                  name: "Request",
                  ofType: null,
                },
              },
            },
            args: [
              {
                name: "account_id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "getRequestsByTargetID",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "LIST",
                ofType: {
                  kind: "OBJECT",
                  name: "Request",
                  ofType: null,
                },
              },
            },
            args: [
              {
                name: "target_account_id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
          {
            name: "getSessionByAccountID",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Session",
                ofType: null,
              },
            },
            args: [
              {
                name: "account_id",
                type: {
                  kind: "NON_NULL",
                  ofType: {
                    kind: "SCALAR",
                    name: "Any",
                  },
                },
              },
            ],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "Request",
        fields: [
          {
            name: "account",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "account_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "status",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "target_account",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "OBJECT",
                name: "Account",
                ofType: null,
              },
            },
            args: [],
          },
          {
            name: "target_account_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: "OBJECT",
        name: "Session",
        fields: [
          {
            name: "account_id",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "session",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
          {
            name: "updatedAt",
            type: {
              kind: "NON_NULL",
              ofType: {
                kind: "SCALAR",
                name: "Any",
              },
            },
            args: [],
          },
        ],
        interfaces: [],
      },
      {
        kind: "SCALAR",
        name: "Any",
      },
    ],
    directives: [],
  },
} as unknown as IntrospectionQuery;
