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
    let rows = lines text
    mapM_ putStrLn rows
    let words = map (!! 0) rows
    putStrLn words
    hClose handle
