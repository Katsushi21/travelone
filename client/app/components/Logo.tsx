import Image from 'next/image';

export const Logo: React.FC = () => {
  return (
    <div>
      <Image
        src="/icon.png"
        height={70}
        width={132}
        alt="Logo"
        className="cursor-pointer object-contain"
      />
    </div>
  );
};
