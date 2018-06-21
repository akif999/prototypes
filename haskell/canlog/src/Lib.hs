module Lib
    ( someFunc
    ) where

import System.IO
import System.Environment
import Data.List

someFunc :: IO ()
someFunc = do
    args <- getArgs
    handle <- openFile (args !! 0) ReadMode
    text <- hGetContents handle
    let tops =  map (!! 0) $ map words $ lines text
    mapM_ putStrLn tops
    hClose handle
