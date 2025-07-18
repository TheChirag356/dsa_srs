import React from "react";
import { Button } from "./ui/button";
import { GithubIcon, Terminal } from "lucide-react";
import Dither from "./Dither";

const Hero = () => {
  return (
    <div className="relative flex h-[85vh] w-full flex-col items-center justify-center gap-8 font-[family-name:var(--font-space-grotesk)]">
      <h1 className="text-7xl font-semibold">dsa_srs</h1>
      <p className="text-md text-muted-foreground px-4 text-center md:w-2/5 md:text-lg">
        A terminal-based spaced repetition tool to master DSA concepts and
        LeetCode problems.
      </p>
      <p className="text-muted-foreground text-sm">
        Built with Go, Bubble Tea, and Lip Gloss
      </p>

      <div className="flex items-center justify-center gap-4">
        <Button asChild>
          <a
            href="https://github.com/TheChirag356/dsa_srs"
            target="_blank"
            rel="noopener noreferrer"
          >
            <GithubIcon className="mr-2 h-4 w-4" />
            View on GitHub
          </a>
        </Button>
        <Button variant="outline" asChild>
          <a
            href="https://pkg.go.dev/github.com/TheChirag356/dsa_srs"
            target="_blank"
            rel="noopener noreferrer"
          >
            <Terminal className="mr-2 h-4 w-4" />
            Install via Go
          </a>
        </Button>
      </div>
      <div style={{ width: "100%", height: "600px", position: "absolute", zIndex: "-10" }} className="opacity-10">
        <Dither
          waveColor={[0.5, 0.5, 0.5]}
          disableAnimation={false}
          enableMouseInteraction={true}
          mouseRadius={0.3}
          colorNum={4}
          waveAmplitude={0.3}
          waveFrequency={3}
          waveSpeed={0.05}
        />
      </div>
    </div>
  );
};

export default Hero;
