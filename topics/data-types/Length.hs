module Length where

newtype Feet = Feet Float

newtype Meters = Meters Float

instance Show Feet where
  show (Feet v) = show v ++ "ft"

instance Show Meters where
  show (Meters v) = show v ++ "m"

class ToFeet a where
    toFeet :: a -> Feet

class ToMeters a where
    toMeters :: a -> Meters

instance ToFeet Feet where
    toFeet v = v

instance ToMeters Feet where
    toMeters (Feet v) = Meters (v * 0.3048)

instance ToFeet Meters where
    toFeet (Meters v) = Feet (v * 3.28084)

instance ToMeters Meters where
    toMeters v = v

addImperial :: (ToFeet a) => (ToFeet b) => a -> b -> Feet
addImperial l0 l1 = Feet (f0 + f1)
    where Feet f0 = toFeet l0
          Feet f1 = toFeet l1

addMetric :: (ToMeters a) => (ToMeters b) => a -> b -> Meters
addMetric l0 l1 = Meters (m0 + m1)
    where Meters m0 = toMeters l0
          Meters m1 = toMeters l1
