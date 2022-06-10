import { BellIcon, SearchIcon } from '@heroicons/react/solid';
import Link from 'next/link';
import React, { useEffect, useState } from 'react';

import { Logo } from './Logo';

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

  return (
    <header className={`${isScrolled && 'bg-black'}`}>
      <div className="flex items-center space-x-2 md:space-x-10">
        <Logo />
        <ul className="hidden space-x-4 md:flex">
          <li className="headerLink">a</li>
          <li className="headerLink">b</li>
          <li className="headerLink">c</li>
          <li className="headerLink">d</li>
          <li className="headerLink">e</li>
        </ul>
      </div>
      <div className="flex item-center space-x-4 text-sm font-light">
        <SearchIcon className="hidden h-6 w-6 sm:inline" />
        <p className="hidden lg:inline">Kids</p>
        <BellIcon className="h-6 w-6" />
        <Link href="/mypage">mypage</Link>
      </div>
    </header>
  );
};
