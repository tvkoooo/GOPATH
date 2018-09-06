
function shard_csv_file_save(directory, prefix, suffix, arr, shard)
	file = nil
	repeat
		if 0 >= table.getn(arr) then
			break
		end
		-- start form 1 not 0.
		local file_index = shard + 1
		local source = directory.."/"..prefix..file_index..suffix
		file = io.open(source,"wb")
		if file == nil then
			print("file:"..source.." is a invalid.")
			break
		end
		for i,v in ipairs(arr) do
			file:write(v.."\n")
		end
		--
		file:close()
	until true
end

function shard_csv_file(directory, filename, shard, prefix, suffix)
	repeat
		file = nil
		if 1 > shard then
			print("shard must big then 1.")
			break
		end
		local source = directory.."/"..filename
		file = io.open(source,"rb")
		if file == nil then
			print("the source file:"..source.." is a invalid.")
			break
		end
		--
		local arrays = {}
		local length = 0
		--
		for line in file:lines() do
			local uid = tonumber(line)
			if 0 ~= uid and nil ~= uid then
				table.insert(arrays, uid)
				length = length + 1
			end
		end
		file:close()
		--
		local shard_number = math.floor( length / shard )
		--
		local shard_arrays = {}
		--
		for i=0,shard - 1,1 do
			shard_arrays[i] = {}
		end
		--
		for i,v in ipairs(arrays) do
			local s = math.floor( v % shard )
			local arr = shard_arrays[s]
			table.insert(arr, v)
		end
		--
		for i=0,shard - 1,1 do
			local arr = shard_arrays[i]
			shard_csv_file_save(directory, prefix, suffix, arr, i)
		end
	until true
end

function shard_robot_csv_file()
	shard_csv_file("1", "robot.csv", 2, "robot_", ".list")
	shard_csv_file("2", "robot.csv", 2, "robot_", ".list")
	shard_csv_file("3", "robot.csv", 4, "robot_", ".list")
	shard_csv_file("9", "robot.csv", 4, "robot_", ".list")
end

shard_robot_csv_file()
