export type loginSchema = {
  email: string;
  password: string;
};

export type userSchema = {
  id: number;
  name: string;
  email: string;
  avatar: string;
  token: string;
  state: string;
};

export type accountsSchema = {
  id: number;
  name: string;
  email: string;
  gender: 'm' | 'f';
  birth: number;
  history: number;
  avatar: string;
  favorite: string;
  introduction: string;
};

export type articleSchema = {
  id: number;
  uid: number;
  title: string;
  thumbnail: string;
  address: string;
  conditions: object;
  comment1: string;
  comment2: string;
  comment3: string;
  image1: string;
  image2: string;
  image3: string;
  modification: string;
  registration: string;
};

export type searchSchema = {
  prefecture: number;
  conditions: object;
  keyword: string;
};
