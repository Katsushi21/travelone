import React from 'react'
import Image from 'next/image'

export const Logo: React.FC = () => {
  return (
    <div>
      <Image
        src='/icon.jpg'
        height={144}
        width={144}
        alt='Logo'
        className='cursor-pointer object-contain'
      />
    </div>
  )
}
