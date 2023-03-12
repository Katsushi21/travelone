'use client';
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
    <div className={`navbar fixed top-0 z-50 duration-[.3s] ${headerColor}`}>
      <div className="navbar-start">
        <Logo />
        <ul className="hidden space-x-8 px-6 md:flex">
          <li className="headerLink">Home</li>
          <li className="headerLink">About</li>
          <li className="headerLink">Posts</li>
          <li className="headerLink">Maps</li>
          <li className="headerLink">Friends</li>
        </ul>
      </div>
      <div className="navbar-end">
        <Searchbar />
        <Bell />
        <Avatar />
      </div>
    </div>
  );
};
