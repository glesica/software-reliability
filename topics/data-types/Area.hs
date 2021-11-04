module Area where

import Length

data SquareFeet = SquareFeet Feet Feet

data SquareMeters = SquareMeters Meters Meters

instance Show SquareFeet where
    show (SquareFeet w h) = show (wf * hf) ++ "ft^2"
        where Feet wf = w
              Feet hf = h

instance Show SquareMeters where
    show (SquareMeters w h) = show (wm * hm) ++ "m^2"
        where Meters wm = w
              Meters hm = h

areaImperial :: (ToFeet a) => (ToFeet b) => a -> b -> SquareFeet
areaImperial w h = SquareFeet wf hf
    where wf = toFeet w
          hf = toFeet h

areaMetric :: (ToMeters a) => (ToMeters b) => a -> b -> SquareMeters
areaMetric w h = SquareMeters wm hm
    where wm = toMeters w
          hm = toMeters h
