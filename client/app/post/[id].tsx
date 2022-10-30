import { GetStaticPaths, GetStaticProps } from 'next';
import Head from 'next/head';
import React from 'react';

type PathParams = {
  id: string;
};

type PageProps = {
  title: string;
  author: string;
};

export const getStaticPaths: GetStaticPaths<PathParams> = async () => {
  return {
    paths: [
      { params: { id: '001' } },
      { params: { id: '002' } },
      { params: { id: '003' } },
    ],
    fallback: false,
  };
};

export const getStaticProps: GetStaticProps<PageProps> = async (context) => {
  const { id } = context.params as PathParams;

  const props: PageProps = {
    title: `Title-${id}`,
    author: `Author-${id}`,
  };

  return { props };
};

const BookPage: React.FC<PageProps> = ({ title, author }: PageProps) => {
  return (
    <>
      <Head>
        <title>{title} の詳細ページ</title>
      </Head>
      <ul>
        <li>タイトル: {title}</li>
        <li>著者: {author}</li>
      </ul>
    </>
  );
};
export default BookPage;
