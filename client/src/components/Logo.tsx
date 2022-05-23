import React from 'react'
import Image from 'next/image'

export const Logo: React.FC = () => {
  return (
    <div>
      <Image
        src='/icon.png'
        height={80}
        width={132}
        alt='Logo'
        className='cursor-pointer object-contain'
      />
    </div>
  )
}
