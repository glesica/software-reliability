module Colors where

import Data.Word

-- | A color represented by red, green, and blue channels.
data RGB = RGB { red :: Word8
               , green :: Word8
               , blue :: Word8
               } deriving (Bounded, Eq, Show)

-- | An RGB color using integral components.
--
-- Examples:
--
-- >>> mix (RGB 0 0 0) (RGB 255 255 255)
-- RGB {red = 127, green = 127, blue = 127}
-- >>> mix (RGB 0 100 200) (RGB 100 0 200)
-- RGB {red = 50, green = 50, blue = 200}
-- >>> mix (RGB 10 20 0) (RGB 20 10 0)
-- RGB {red = 15, green = 15, blue = 0}
mix :: RGB -> RGB -> RGB
mix c0 c1 = RGB rr gg bb
    where addChan c0 c1 = d + m
              where d = c0 `div` 2 + c1 `div` 2
                    m = (c0 `mod` 2 + c1 `mod` 2) `div` 2
          rr = addChan (red c0) (red c1)
          gg = addChan (green c0) (green c1)
          bb = addChan (blue c0) (blue c1)
