module LuciansLusciousLasagna

let expectedMinutesInOven = 40

let remainingMinutesInOven cookTime = expectedMinutesInOven - cookTime

let preparationTimeInMinutes layers = 2 * layers

let elapsedTimeInMinutes layers cookTime =
    preparationTimeInMinutes layers + cookTime

