import Link from 'next/link';
import React, { useEffect, useState } from 'react';
import { Avatar } from './Avatar';
import { Bell } from './Bell';

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

  const headerColor = isScrolled ? 'bg-base-100' : 'bg-opacity-0';

  return (
    <div className={headerColor}>
      <div className={`navbar duration-[.3s] fixed top-0 z-50 ${headerColor}`}>
        <div className="flex-1">
          <Logo />
          <ul className="hidden space-x-8 md:flex px-6">
            <li className="headerLink">Home</li>
            <li className="headerLink">About</li>
            <li className="headerLink">Posts</li>
            <li className="headerLink">Maps</li>
            <li className="headerLink">Friends</li>
          </ul>
        </div>
        <div className="flex-none gap-2">
          <Searchbar />
          <Bell />
          <Avatar />
        </div>
      </div>
    </div>
  );
};
