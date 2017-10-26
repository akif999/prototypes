import System.IO
import System.Exit
import System.Environment(getArgs)
import Data.Text.IO as DTI

data Srec = Srec { srectype :: Int, length :: Int, address :: Int, bytes :: [Int], checksum :: Int } deriving Show

get_and_print :: Handle -> IO ()
get_and_print fp = do
    eof <- hIsEOF fp
    if eof then return ()
    else do
        str <- DTI.hGetLine fp
        print str
        get_and_print fp

usage = print $ "Usage : input filename"

main :: IO ()
main = do
    args <- getArgs
    if null args
        then do
            usage
            exitSuccess
        else do
            print head[args]
            --fp <- openFile filename ReadMode
            --get_and_print fp
