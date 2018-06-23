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
    let heads =  extractFields $ extractRecords text
    mapM_ putStrLn heads
    hClose handle

extractRecords :: String -> [String]
extractRecords rs = lines rs

extractFields :: [String] -> [String]
extractFields fs = map (!! 0) $ map words fs
