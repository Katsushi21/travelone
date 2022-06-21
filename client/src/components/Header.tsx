import { BellIcon, SearchIcon } from '@heroicons/react/solid';
import Link from 'next/link';
import React, { useEffect, useState } from 'react';
import { Avatar } from './Avatar';

import { Logo } from './Logo';
import { Searchbar } from './Searchbar';

export const Header: React.FC = () => {
  const [isScrolled, setIsScrolled] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      if (window.scrollY > 0) {
        setIsScrolled(true);
      } else {
        setIsScrolled(false);
      }
    };
    window.addEventListener('scroll', handleScroll);

    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, []);

  const headerColor = isScrolled ? 'bg-white' : '';

  return (
    // <header className={headerColor}>
    //   <div className="flex items-center space-x-2 md:space-x-8">
    //     <Logo />
    //     <ul className="hidden space-x-8 md:flex">
    //       <li className="headerLink">Home</li>
    //       <li className="headerLink">About</li>
    //       <li className="headerLink">Posts</li>
    //       <li className="headerLink">Maps</li>
    //       <li className="headerLink">Friends</li>
    //     </ul>
    //   </div>
    //   <div className="flex item-center space-x-4 text-sm font-light">
    //     <SearchIcon className="hidden h-6 w-6 sm:inline" />
    //     <BellIcon className="h-6 w-6" />
    //     <Link href="/mypage">Mypage</Link>
    //   </div>
    // </header>
    <div className={headerColor}>
      <div className="navbar bg-base-100">
        <div className="flex-1">
          <Logo />
          <ul className="hidden space-x-8 md:flex">
            <li className="headerLink">Home</li>
            <li className="headerLink">About</li>
            <li className="headerLink">Posts</li>
            <li className="headerLink">Maps</li>
            <li className="headerLink">Friends</li>
          </ul>
        </div>
        <div className="flex-none gap-2">
          <Searchbar />
          <BellIcon className="h-6 w-6" />
          <Avatar />
        </div>
      </div>
    </div>
  );
};
