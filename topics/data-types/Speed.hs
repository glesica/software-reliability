module Speed where

import Length

newtype Seconds = Seconds Float

newtype FeetPerSec = FeetPerSec Feet

-- TODO: Implement MetersPerSec

instance Show FeetPerSec where
    show (FeetPerSec v) = show v ++ "/s"

-- TODO: Implement Show for MetersPerSec

-- | Return the speed necessary to move the given distance,
-- | in feet, in the given number of seconds.
--
-- >>> speedImperial (Feet 10) (Seconds 2)
-- 5.0ft/s
speedImperial :: (ToFeet a) => a -> Seconds -> FeetPerSec
speedImperial d (Seconds t) = FeetPerSec $ Feet (df / t)
    where Feet df = toFeet d

-- | Return the speed necessary to move the given distance,
-- | in meters, in the given number of seconds.
--
-- >>> speedMetric (Meters 10) (Seconds 2)
-- 5.0m/s
--speedMetric :: (ToMeters a) => a -> Seconds -> MetersPerSec
-- TODO: Implement speedMetric

-- | Return the number of feet traveled at the given speed,
-- | in the given number of seconds.
--
-- >>> feetAfter (FeetPerSec $ Feet 5) (Seconds 2)
-- 10.0ft
--feetAfter :: FeetPerSec -> Seconds -> Feet
-- TODO: Implement feetAfter

-- | Return the number of meters traveled at the given speed,
-- | in the given number of seconds.
--
-- >>> metersAfter (MetersPerSec $ Meters 5) (Seconds 2)
-- 10.0m
--metersAfter :: MetersPerSec -> Seconds -> Meters
-- TODO: Implement metersAfter
