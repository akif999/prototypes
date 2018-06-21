import System.IO
import System.Environment
import Data.List

main :: IO()
main = do
    args <- getArgs
    handle <- openFile (args !! 0) ReadMode
    text <- hGetContents handle
    rows <- splitOn "\n" text
    putStr rows
    hClose handle
