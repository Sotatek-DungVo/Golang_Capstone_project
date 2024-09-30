import { useState } from "react";
import { PlayerInfoDetail, PlayerListParams } from "../../services/player/player.type";
import { PlayerService } from "../../services/player/player.service";
import PlayerList from "./PlayerList";

type SearchPlayerProps = {
    fetchPlayer: (params: PlayerListParams) => Promise<void>
}

const SearchPlayer: React.FC<SearchPlayerProps> = ({
    fetchPlayer,
}) => {
    const [username, setUsername] = useState<string | undefined>(undefined)
    const [gender, setGender] = useState<string | undefined>(undefined)
    const [searchPlayers, setSearchPlayers] = useState<PlayerInfoDetail[] | null>(null);
    const [page, setPage] = useState<number>(1);
    const [pageSize, setPageSize] = useState<number>(20);
    const [loading, setIsLoading] = useState<boolean>(false);

    const handleClickSearch = async () => {
        try {
            setIsLoading(true)
            setPage(1);

            const data = await PlayerService.all({
                gender, username, page, pageSize
            })


            setSearchPlayers(data)
            setIsLoading(false)
        } catch (error) {
            console.log('error: ', error)
            setIsLoading(false)
        }
    }

    return (
        <>
        <div className="flex justify-between px-10 pt-5">
            {/* Filter options */}
            <div className="flex gap-x-2">
            <select className="px-2 py-1 border-2 rounded-full" onChange={(e: React.ChangeEvent<HTMLSelectElement>) => {
                setGender(e.target.value)
            }}
                value={gender}
            >
                <option value="">Gender</option>
                <option value="FEMALE">Female</option>
                <option value="MALE">Male</option>
            </select>

            <input
                className="px-2 border-2 rounded-full"
                placeholder="username"
                value={username}
                onChange={(e) => {
                    setUsername(e.target.value)
                }}
            />

            </div>
           
            {/* Search btn */}
            <button
            onClick={handleClickSearch}
            className="min-w-20 px-2 font-semibold text-white rounded-full bg-[#f0564a]">
                Search
            </button>
        </div>

        {searchPlayers &&
        (
            <>
               <h3 className="px-10 pt-10 font-bold text-red-500">Search result</h3>

                <PlayerList playerData={searchPlayers} />
            </>
   )
       }
        </>
    )
}

export default SearchPlayer