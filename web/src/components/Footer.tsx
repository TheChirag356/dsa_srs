import { Twitter } from "lucide-react";
import Link from "next/link";
import React from "react";

const Footer = () => {
  return (
    <div className="flex h-[3rem] w-full items-center justify-center font-[family-name:var(--font-space-grotesk)]">
      Made with love by{"  "}
      <Link
        href={"https://x.com/chiragkun"}
        target="_blank"
        className="mx-1 inline-flex items-center hover:underline"
      >
        <Twitter className="aspect-square h-5" /> @chiragkun
      </Link>
    </div>
  );
};

export default Footer;
