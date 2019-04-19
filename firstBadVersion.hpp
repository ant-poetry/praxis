// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
    int currentBadVersion = -1;
public:
    int firstBadVersion(int n) {
        if(currentBadVersion == -1)
        {
            //错误的
            int left = 1;
            int right = n + 1;
            while(left < right){
                int mid = (left + right) / 2;
                if(isBadVersion(mid)){ //找左边
                    right = mid ;
                }else{
                    left = mid + 1;
                }
            }
            currentBadVersion = left;
            return currentBadVersion;
        }else{
            if(n >= currentBadVersion){
                return n;
            }else{
                return currentBadVersion;
            }
        }
    }
};