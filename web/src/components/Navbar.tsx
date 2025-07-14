import React from "react";
import { Button } from "./ui/button";
import { GithubIcon, Star } from "lucide-react";
import Link from "next/link";

const Navbar = () => {
  return (
    <nav className="font-[family-name:var(--font-space-grotesk)] sticky bg-background z-50 top-0 left-0 flex h-[3.5rem] w-full items-center justify-between gap-6 border-b-[1px] border-neutral-600 px-4 py-2">
      <div>dsa_srs</div>
      <div className="flex items-center justify-evenly gap-6">
        <Button variant="ghost" className="cursor-default">
          <Star />
          Stars
        </Button>
        <Link href={"https://github.com/TheChirag356/dsa_srs"} target="_blank">
          <Button>
            <GithubIcon />
            Github
          </Button>
        </Link>
      </div>
    </nav>
  );
};

export default Navbar;
