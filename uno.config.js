import { defineConfig, presetWind3, presetAttributify, presetWebFonts } from 'unocss'

const range = (n, fn) => Array.from({ length: n }, (_, i) => fn(i));

function rangeWithPrefixes(n, base, prefixes = ['']) {
  return prefixes.flatMap(prefix =>
    range(n, i => `${prefix}${base}-${i}`)
  );
}

export default defineConfig({
  content: {
    filesystem: ['./**/*.{html,md,js}'],
  },
  presets: [
    presetAttributify(),
    presetWind3({
      dark: 'media'
    })
  ],
  theme: {
    colors: {
      blockBg: 'var(--block-background-color)',
      bg: 'var(--background-color)',
      textColor: 'var(--text-color)',
      secondary: 'var(--secondary-text-color)'
    },
    fontFamily: {
      mono: ["Courier New", "Courier", "Menlo", "Monaco", "Consolas", "Liberation Mono", "Lucida Console"],
      serif: ["Noto Serif", "Noto Serif SC", "serif"],
    },
    animation: {
      kayframes: {
        fadeIn: '{from { opacity: 0; transform: translateY(1rem); } to { opacity: 1;transform: translateY(0); }}'
      },
      counts: {
        infinite: 'infinite'
      }
    }
  },
  shortcuts: {
    'small-button': 'bg-blockBg px-2 py-1 no-underline hover:no-underline rounded border border-solid border-bg hover:border-textColor transition',
    'round-sticker': 'bg-blockBg rounded-full shadow-sm p-4'
  },
  safelist: [
    'text-sm', 'text-md', 'text-lg', 'text-xl', 'text-2xl', 'text-3xl', 'text-4xl', 'text-5xl',
    'text-base!', 'md:text-size-lg', 'md:text-size-xl', 'md:text-size-2xl',

    'animate-heart-beat', 'animate-fade-in',
    ...'absolute bottom-full left-1/2 transform -translate-x-1/2 translate-y-1 bg-black/80 text-white px-2.5 py-1.5 rounded text-xs whitespace-nowrap opacity-0 invisible transition-all duration-300 pointer-events-none mb-2 z-1000 group-hover:opacity-100 group-hover:visible group-hover:translate-y-0 dark:bg-white/90 dark:text-black'.split(' '),

    ...range(10, i => `opacity-${i * 10}`),
    ...rangeWithPrefixes(5, 'rotate', ['', '-']),
    ...rangeWithPrefixes(5, 'translate-x', ['', '-']),
    ...rangeWithPrefixes(5, 'translate-y', ['', '-']),
    ...rangeWithPrefixes(30, 'w', ['', 'md:', 'lg:']),
    ...rangeWithPrefixes(30, 'h', ['', 'md:']),
    ...rangeWithPrefixes(30, 'right', ['', 'md:', '-', 'md:-']),
    ...rangeWithPrefixes(30, 'top', ['', '-']),

    // special cases
    'absolute',
    'w-[calc(25%_-_1rem)]',
    'w-[calc(50%_-_2rem)]',
    'w-[10rem]',
    'sm:w-full',
  ]
})
