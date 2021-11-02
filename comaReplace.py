import os

cur_dir = os.getcwd()

file_list = []
dir_list = []

for stuff in os.listdir():
    if os.path.isfile(stuff):
        file_list.append(cur_dir + '\\' + stuff)
    else:
        dir_list.append(cur_dir + '\\' + stuff)


while len(dir_list) != 0:
    cur_dir = dir_list.pop(0)  # popstuff

    for stuff in os.listdir(cur_dir):
        path = cur_dir + '\\' + stuff
        if os.path.isdir(path):
            dir_list.append(path)
        else:
            file_list.append(path)

for f in file_list:
    os.rename(f, f.replace(',', '_'))

print("File Renaming Complete")